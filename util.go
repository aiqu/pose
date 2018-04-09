package pose

import (
	"errors"
	"fmt"
	"math"
	"strconv"
	"strings"
)

func NewPose(str string) (*Pose, error) {
	tok := strings.Split(str, ",")
	if len(tok) != 7 {
		return nil, errors.New(fmt.Sprintf("Expect 7 comma-seperated float values for pose, got %s", str))
	}
	pointStr := strings.Join(tok[:3], ",")
	quatStr := strings.Join(tok[3:], ",")
	pose := &Pose{}
	var err error
	pose.Translation, err = NewPoint3(pointStr)
	if err != nil {
		return nil, err
	}
	pose.Rotation, err = NewQuaternion(quatStr)
	if err != nil {
		return nil, err
	}
	return pose, nil
}

func NewPoint3(str string) (*Point3, error) {
	tok := strings.Split(str, ",")
	if len(tok) != 3 {
		return nil, errors.New(fmt.Sprintf("Expect 3 comma-seperated float values for point, got %s", str))
	}
	val := make([]float32, 3)
	for i, s := range tok {
		if v, err := strconv.ParseFloat(s, 32); err != nil {
			return nil, err
		} else {
			val[i] = float32(v)
		}
	}
	return &Point3{X: val[0], Y: val[1], Z: val[2]}, nil
}

func NewQuaternion(str string) (*Quaternion, error) {
	tok := strings.Split(str, ",")
	if len(tok) != 4 {
		return nil, errors.New(fmt.Sprintf("Expect 4 comma-seperated float values for quaternion, got %s", str))
	}
	val := make([]float32, 4)
	for i, s := range tok {
		if v, err := strconv.ParseFloat(s, 32); err != nil {
			return nil, err
		} else {
			val[i] = float32(v)
		}
	}
	return &Quaternion{Qw: val[0], Qx: val[1], Qy: val[2], Qz: val[3]}, nil
}

func (point Point3) Norm() float32 {
	v := point.X * point.X
	v += point.Y * point.Y
	v += point.Z * point.Z
	return float32(math.Sqrt(float64(v)))
}
