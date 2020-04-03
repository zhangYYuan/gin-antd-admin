import React, { useState, useEffect } from "react";
import {  Cascader} from "antd";
import { queryProvince, queryCity, queryCountry } from '@/services/common';



const AddressForm = () =>{
  const [options, setOptions] = useState([])
  const [select, setSelect] = useState([])
  const initProvince = async () => {
    const res = await queryProvince()
    const v = res.resultBody.map(i => ({ label: i.name, value: i.id, level: i.level, isLeaf: false }))
    return v
  }


  const initCity = async (params) => {
    const res = await queryCity(params)
    const v = res.resultBody.map(i => ({ label: i.name, value: i.id, level: i.level, isLeaf: params.provinceById === 1 }))
    return v
  }

  const initCountry = async (params) => {
    const res = await queryCountry(params)
    const v = res.resultBody.map(i => ({ label: i.name, value: i.id, level: i.level, isLeaf: params.provinceById === 1 }))
    return v
  }

  const loadData = async (selectedOptions) => {
    const targetOption = selectedOptions[selectedOptions.length - 1];
    targetOption.loading = true;
    if (targetOption.level === 0) {
      const params = { provinceById: targetOption.value }
      const v = await initCity(params)
      targetOption.loading = false;
      targetOption.children = v
      setOptions([...options])
    } else if (targetOption.level === 1) {
      console.log('loadData', select)
      // const params = { cityId: select[] }
      // const v = await initCountry(params)

    }
  }
  useEffect(  () => {

    initProvince().then(op => {
      setOptions(op)
    })
  }, [])

  const onChange = (value, selectedOptions)  => {
    setSelect(selectedOptions)

  }
  return (
    <Cascader
      options={options}
      loadData={loadData}
      onChange={onChange}
      changeOnSelect
    />
  )
}


export default AddressForm
