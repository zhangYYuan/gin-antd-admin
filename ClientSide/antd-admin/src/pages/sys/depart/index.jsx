import React, { useRef, useState, useEffect } from 'react';
import { DownOutlined, PlusOutlined } from '@ant-design/icons';
import { findDOMNode } from 'react-dom';
import { PageHeaderWrapper } from '@ant-design/pro-layout';
import { connect } from 'umi';

export const BasicList = props => {

  const {
    loading,
    dispatch,
    sysDepart: { list },
  } = props;
  useEffect(() => {
    console.log('---->>>>', props)

    dispatch({
      type: 'sysDepart/fetch',
      payload: {
        count: 5,
      },
    });
  }, []);

  return (
    <div>1222</div>
  )
}


export default connect(({ sysDepart }) => ({
  sysDepart
}))(BasicList);
