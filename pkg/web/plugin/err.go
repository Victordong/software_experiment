package plugin

type CustomErr struct {
	Code        int
	StatusCode  int
	Information string
}

func (err CustomErr) Error() string {
	return err.Information
}
