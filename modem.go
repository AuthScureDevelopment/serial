package serial

import "time"

var (
	MyNumber = "6282325600996"
)

func (s *SerialPort) Call(number string) error {
	if err := s.Println("ATH"); err != nil {
		return err
	}
	if err := s.Println("ATD+" + number + ";"); err != nil {
		return err
	}
	ticker := time.NewTicker(time.Second * 1)
	time.Sleep(time.Second * 9)
	ticker.Stop()
	return s.HangUp()
}

func (s *SerialPort) HangUp() error {
	if err := s.Println("ATH"); err != nil {
		return err
	}
	return nil
}
