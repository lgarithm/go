package monitoring

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/NVIDIA/go-dcgm/pkg/dcgm"
)

type Monitor struct {
	gpus     []uint
	cleanup  func()
	stop     []chan bool
	finished []chan bool
	csvFile  []*os.File
	fields   []dcgm.Short
	Period   time.Duration
}

var defaultFields = []dcgm.Short{
	dcgm.DCGM_FI_PROF_GR_ENGINE_ACTIVE,
	dcgm.DCGM_FI_PROF_SM_ACTIVE,
	dcgm.DCGM_FI_PROF_SM_OCCUPANCY,
	dcgm.DCGM_FI_PROF_PIPE_TENSOR_ACTIVE,
	dcgm.DCGM_FI_PROF_DRAM_ACTIVE,
}

func NewMonitor(period time.Duration) *Monitor {
	return &Monitor{
		fields: defaultFields,
		Period: period,
	}
}

func outfile(i uint) string {
	dir := `log`
	os.MkdirAll(dir, os.ModePerm)
	return fmt.Sprintf("%s/monitoring_%d.csv", dir, i)
}

func (m *Monitor) createCSV() {
	m.csvFile = make([]*os.File, len(m.gpus))

	for _, gpu := range m.gpus {
		// Create a new file and write the header
		file, err := os.Create(outfile(gpu))
		if err != nil {
			log.Panicln(err)
		}

		s := "gpu_id,time,gr_engine_active,sm_active,sm_occupancy,pipe_tensor_active,dram_active\n"
		_, err = file.WriteString(s)
		if err != nil {
			log.Panicln(err)
		}
		m.csvFile[gpu] = file
	}
}

func (m *Monitor) MonitorStatus() {
	go func() {
		for {
			select {
			case <-m.stop[0]:
				m.finished[0] <- true
				return
			default:
				for _, gpu := range m.gpus {
					st, err := dcgm.GetDeviceStatus(gpu)
					if err != nil {
						log.Panicln(err)
					}
					t := time.Now().UnixNano()
					_, err = m.csvFile[0].WriteString(fmt.Sprintf("%d,%d,%d\n", gpu, t, st.Utilization.GPU))
					if err != nil {
						log.Panicln(err)
					}
				}
			}
		}
	}()
}

func (m *Monitor) Start() {
	cleanup, err := dcgm.Init(dcgm.Embedded)
	if err != nil {
		log.Panicln(err)
	}
	m.cleanup = cleanup

	m.gpus, err = dcgm.GetSupportedDevices()
	if err != nil {
		log.Panicln(err)
	}
	log.Printf("GPUs: %q", m.gpus)
	const fieldGroupName = "PROF_ACTIVE"
	fieldsGroup, err := dcgm.FieldGroupCreate(fieldGroupName, m.fields)
	if err != nil {
		log.Panicln(err)
	}

	group, err := dcgm.NewDefaultGroup(fieldGroupName)
	if err != nil {
		log.Panicln(err)
	}

	if err := dcgm.WatchFieldsWithGroupEx(fieldsGroup, group, 1000, 1, 1); err != nil {
		log.Panicln(err)
	}

	// Initialize channels
	m.stop = make([]chan bool, len(m.gpus))
	m.finished = make([]chan bool, len(m.gpus))
	for i := range m.gpus {
		m.stop[i] = make(chan bool)
		m.finished[i] = make(chan bool)
	}

	m.createCSV()

	for _, gpu := range m.gpus {
		gpu := gpu
		go m.MonitorSMACT(gpu)
	}
}

func (m *Monitor) Stop() {
	for _, gpu := range m.gpus {
		m.stop[gpu] <- true
		<-m.finished[gpu]
		m.csvFile[gpu].Close()
	}

	m.cleanup()
}

func (m *Monitor) MonitorSMACT(gpu uint) {
	for {
		select {
		case <-m.stop[gpu]:
			m.finished[gpu] <- true
			return
		default:
			values, err := dcgm.GetLatestValuesForFields(gpu, m.fields)
			if err != nil {
				log.Panicln(err)
			}

			gr := values[0].Float64()
			sm := values[1].Float64()
			occ := values[2].Float64()
			pipe := values[3].Float64()
			dram := values[4].Float64()
			t := time.Now().UnixNano()
			s := fmt.Sprintf("%d,%d,%f,%f,%f,%f,%f\n", gpu, t, gr, sm, occ, pipe, dram)
			fmt.Fprint(os.Stderr, s)
			if gr != 0 || sm != 0 || occ != 0 || pipe != 0 || dram != 0 {
				// fmt.Fprint(os.Stderr, s)
			} else {
				// fmt.Fprintf(os.Stderr, "all zero!\n")
			}
			_, err = m.csvFile[gpu].WriteString(s)
			if err != nil {
				log.Panicln(err)
			}
		}
		time.Sleep(m.Period)
	}
}
