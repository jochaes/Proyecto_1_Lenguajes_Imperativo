package data_structures

import (
	"errors"
	"fmt"
)

//Tipo para el Almacen
type MapTable map[string]any

/*
Inserta un key:value / variable:dato en el almacen 
	key: nombre de la variable
	value: dato que "guarda" la variable, puede ser de cualquier tipo 
*/
func (storage *MapTable) Insert(key string, value any) error{

	if (*storage)[key] == nil {
		(*storage)[key] = value
		fmt.Println("  mapTable(Insert)", key, value)
		return nil
	}else{
		return errors.New("mapTable(insert) error: key is already in storage")
	}
}

/*
Actualiza una variable en el almacen
	key: nombre de la variable
	value: nuevo dato que va a guardar la variable, puede ser de cualquier tipo 
*/
func (storage *MapTable) Update(key string, value any) error{

	if (*storage)[key] != nil {
		(*storage)[key] = value
		fmt.Println("  mapTable(Update)", key, "to", value)
		return nil
	}else{
		return errors.New("mapTable(update) error: Key is not in storage")
	}
}

/*
Devuelve el dato que guarda la variable que se guarda en el almacen 
key: nombre de la variable
*/
func (storage *MapTable) Get(key string) (any, error){

	if val := (*storage)[key]; val != nil {
		fmt.Println("  mapTable(Get)", key, val)
		return val, nil
	}else{
		return nil, errors.New("mapTable(get) error: Key is not in storage")
	}
}

/*
Borra una variable del almacen
 key: Nombre de la variable
*/
func (storage *MapTable) Delete(key string) error{

	if val := (*storage)[key]; val != nil {
		delete(*storage, key)
		fmt.Println("  mapTable(Delete)", key, val)
		return nil
	}else{
		return errors.New("mapTable(delete) error: Key is not in storage")
	}	
}
