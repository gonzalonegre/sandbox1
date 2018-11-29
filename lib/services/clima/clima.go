package clima

import (
	"fmt"
	"strings"
	"strconv"
	"golang.org/x/net/context"
	"google.golang.org/appengine/datastore"
	"google.golang.org/appengine/log"
)

const (
	index           = `Dias`
	invalidData = `error: no disponemos de datos`
)

// defines attributes
type Clima struct {
	Dia  int `json:"dia"`
	Estado string `json:"estado"`
}


// Create an clima
func Create(c context.Context, clm *Clima) (*Clima, error) {
	var output *Clima
	if clm == nil {
		return nil, fmt.Errorf(invalidData)
	}

	output, _ = GetByDia(c, clm.Dia)

	if output == nil {
		key := datastore.NewKey(c, index, strconv.Itoa(clm.Dia+1), 0, nil)
		insKey, err := datastore.Put(c, key, clm)

		if err != nil {
			log.Errorf(c, "ERROR INSERTING DIA: %v", err.Error())
			return nil, err
		}

		value,_ :=strconv.Atoi(insKey.StringID())
		output, err = GetByDia(c, value)
		if err != nil {
			log.Errorf(c, "ERROR GETTING DIA OUTPUT: %v", err.Error())
			return nil, err
		}
		return output, nil
	}
	log.Infof(c, "Clima was previously saved: %v", clm.Dia)
	return output, nil
}

// GetByDia an clima based on its Dia
func GetByDia(c context.Context, dia int) (*Clima, error) {
	
	storedia:=dia%360
	if(storedia<0){storedia+=360}
	key := datastore.NewKey(c, index, strconv.Itoa(storedia+1), 0, nil)
	var clm Clima
	err := datastore.Get(c, key, &clm)

	if err != nil {
		if strings.HasPrefix(err.Error(), `datastore: no such entity`) {
			err = fmt.Errorf(invalidData, dia)
		}
		return nil, err
	}
	clm.Dia = dia
	return &clm, nil
}

// GetDias Fetches all dias
func GetDias(c context.Context) ([]Clima, error) {
	var output []Clima
	q := datastore.NewQuery(index)
	_, err := q.GetAll(c, &output)

	if err != nil {
		log.Errorf(c, "error recuperando dias")
		return nil, err
	}

	if len(output) <= 0 {
		return nil, fmt.Errorf("no se encontraron dias")
	}
	return output, nil
}

// Update clima data
func Update(c context.Context, clm *Clima) (*Clima, error) {
	if clm == nil {
		return nil, fmt.Errorf(invalidData)
	}

	output, _ := GetByDia(c, clm.Dia)
	if output != nil {
		key := datastore.NewKey(c, index, strconv.Itoa(clm.Dia+1), 0, nil)
		insKey, err := datastore.Put(c, key, clm)

		if err != nil {
			log.Errorf(c, "ERROR UPDATING DIA: %v", err.Error())
			return nil, err
		}

		value,_ :=strconv.Atoi(insKey.StringID())
		output, err = GetByDia(c, value)
		if err != nil {
			log.Errorf(c, "ERROR GETTING DIA OUTPUT: %v", err.Error())
			return nil, err
		}
		return output, nil
	}
	return nil, fmt.Errorf(`Dia '%v' not found`, clm.Dia)
}

// Delete an clima based on its Dia.
func Delete(c context.Context, dia int) error {
	var output *Clima
	output, _ = GetByDia(c, dia)

	if output != nil {
		log.Infof(c, "Deleting clima, dia: %v", dia)
		key := datastore.NewKey(c, index, strconv.Itoa(dia+1), 0, nil)
		err := datastore.Delete(c, key)

		if err != nil {
			log.Errorf(c, "ERROR DELETING DIA: %v", err.Error())
			return err
		}
		return nil
	}
	return fmt.Errorf("clima '%v' don't exist on the database", dia)
}