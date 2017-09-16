package air_test

import (
    "encoding/json"
    "fmt"
    "github.com/vdods/air"
    "github.com/stretchr/testify/assert"
    "testing"
)

func TestStuff (t *testing.T) {
    // Ensure Separator is set to the expected value
    assert.Equal(t, "\n", air.GetSeparator())
    assert.Equal(t, air.DEFAULT_SEPARATOR, air.GetSeparator())

    // If nil is passed in for the original error, then just start with the given formatted error message.
    // Equivalently, one can call `air.Roar(nil, format, values...)`.
    err := air.Roar(nil, "HIPPO %d", 123)
    assert.Equal(t, "HIPPO 123", err.Error())
    err = air.Roar(err, "OSTRICH %d", 456)
    assert.Equal(t, "HIPPO 123\nOSTRICH 456", err.Error())

    // Set Separator to something else and run the same tests
    air.SetSeparator(";;;")
    assert.Equal(t, ";;;", air.GetSeparator())
    err = air.Roar(nil, "HIPPO %d", 123)
    assert.Equal(t, "HIPPO 123", err.Error())
    err = air.Roar(err, "OSTRICH %d", 456)
    assert.Equal(t, "HIPPO 123;;;OSTRICH 456", err.Error())

    // Set Separator back to the default.
    air.SetSeparator(air.DEFAULT_SEPARATOR)
    assert.Equal(t, air.DEFAULT_SEPARATOR, air.GetSeparator())
}

func TestMoreStuff (t *testing.T) {
    // air.Roar can accept any instance of the error interface as the starting error.
    err := fmt.Errorf("THINGY %d", 999)
    assert.Equal(t, "THINGY 999", err.Error())
    err = air.Roar(err, "OSTRICH %d", 456)
    assert.Equal(t, "THINGY 999\nOSTRICH 456", err.Error())
}

func TestMoreStuff2 (t *testing.T) {
    // air.Errorf is a drop-in replacement to fmt.Errorf, except that it
    // returns *AirRoar (which implements the error interface)
    err := air.Errorf("THINGY %d", 999)
    assert.Equal(t, "THINGY 999", err.Error())
    err = air.Roar(err, "OSTRICH %d", 456)
    assert.Equal(t, "THINGY 999\nOSTRICH 456", err.Error())
}

func TestYetMoreStuff (t *testing.T) {
    err := air.Roar(nil, "HIPPO %d", 123)
    assert.Equal(t, "HIPPO 123", err.Error())
    err2 := air.Roar(err, "OSTRICH %d", 456)
    assert.Equal(t, "HIPPO 123", err.Error())
    assert.Equal(t, "HIPPO 123\nOSTRICH 456", err2.Error())
}

func Divide (num, den float64) (float64, error) {
    if den == 0 {
        return 0, fmt.Errorf("divide by zero")
    }
    return num/den, nil
}

func Mean (values []float64) (float64, error) {
    var sum float64 = 0
    for _,value := range values {
        sum += value
    }
    var average float64
    var err error
    average,err = Divide(sum,float64(len(values)))
    if err != nil {
        return 0, air.Roar(err, "in Mean (len(values) = %d)", len(values))
    }
    return average, nil
}

func MeanOfJSONValues (json_string string) (float64, error) {
    var values []float64
    err := json.Unmarshal([]byte(json_string), &values)
    if err != nil {
        return 0, air.Roar(err, "in MeanOfJSONValues")
    }
    mean,err := Mean(values)
    if err != nil {
        return 0, air.Roar(err, "in MeanOfJSONValues")
    }
    return mean, nil
}

func TestMoreRealisticExample (t *testing.T) {
    var err error

    _,err = MeanOfJSONValues("[]")
    assert.Equal(t, "divide by zero\nin Mean (len(values) = 0)\nin MeanOfJSONValues", err.Error())

    _,err = MeanOfJSONValues("&*(")
    assert.Equal(t, "invalid character '&' looking for beginning of value\nin MeanOfJSONValues", err.Error())

    mean,err := MeanOfJSONValues("[1, 1, 2, 4]")
    assert.Equal(t, float64(2), mean)
    assert.Nil(t, err)
}
