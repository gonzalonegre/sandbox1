package clima

import (
	"fmt"
	"strings"

	"golang.org/x/net/context"
	"google.golang.org/appengine/datastore"
	"google.golang.org/appengine/log"
)

const (
	index           = `Dia`
	invalidUserData = `error: no disponemos de datos`
)

// User defines attributes
type Clima struct {
	Dia  string `json:"dia"`
	Estado string `json:"estado"`
}


// Create an clima
func Create(c context.Context, clm *Clima) (*Clima, error) {
	var output *Clima
	if clm == nil || clm.Dia == `` {
		return nil, fmt.Errorf(invalidUserData)
	}

	output, _ = GetByDia(c, clm.Dia)

	if output == nil {
		key := datastore.NewKey(c, index, clm.Dia, 0, nil)
		insKey, err := datastore.Put(c, key, clm)

		if err != nil {
			log.Errorf(c, "ERROR INSERTING DIA: %v", err.Error())
			return nil, err
		}

		output, err = GetByDia(c, insKey.StringID())
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
func GetByDia(c context.Context, dia string) (*Clima, error) {
	if dia == `` {
		return nil, fmt.Errorf(invalidUserData)
	}
	userKey := datastore.NewKey(c, index, dia, 0, nil)
	var clm Clima
	err := datastore.Get(c, userKey, &clm)

	if err != nil {
		if strings.HasPrefix(err.Error(), `datastore: no such entity`) {
			err = fmt.Errorf(invalidUserData, dia)
		}
		return nil, err
	}
	return &clm, nil
}

// GetDias Fetches all users
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
	if clm == nil || clm.Dia == `` {
		return nil, fmt.Errorf(invalidUserData)
	}

	output, _ := GetByDia(c, clm.Dia)
	if output != nil {
		key := datastore.NewKey(c, index, clm.Dia, 0, nil)
		insKey, err := datastore.Put(c, key, clm)

		if err != nil {
			log.Errorf(c, "ERROR UPDATING DIA: %v", err.Error())
			return nil, err
		}

		output, err = GetByDia(c, insKey.StringID())
		if err != nil {
			log.Errorf(c, "ERROR GETTING DIA OUTPUT: %v", err.Error())
			return nil, err
		}
		return output, nil
	}
	return nil, fmt.Errorf(`Dia '%v' not found`, clm.Dia)
}

// Delete an clima based on its Dia.
func Delete(c context.Context, dia string) error {
	var output *Clima
	output, _ = GetByDia(c, dia)

	if output != nil {
		log.Infof(c, "Deleting clima, dia: %v", dia)
		key := datastore.NewKey(c, index, dia, 0, nil)
		err := datastore.Delete(c, key)

		if err != nil {
			log.Errorf(c, "ERROR DELETING DIA: %v", err.Error())
			return err
		}
		return nil
	}
	return fmt.Errorf("clima '%v' don't exist on the database", dia)
}