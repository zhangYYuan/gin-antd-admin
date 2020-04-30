import React, { useRef, useState, useEffect } from 'react';
import { DownOutlined, PlusOutlined } from '@ant-design/icons';
import { findDOMNode } from 'react-dom';
import { PageHeaderWrapper } from '@ant-design/pro-layout';
import { connect } from 'umi';

export const BasicList = props => {

  const {
    loading,
    dispatch,
    listAndbasicList: { list },
  } = props;
  useEffect(() => {
    dispatch({
      type: 'listAndbasicList/fetch',
      payload: {
        count: 5,
      },
    });
  }, []);

  return (
    <div>1222</div>
  )
}


export default connect(({ listAndbasicList, loading }) => ({
  listAndbasicList
}))(BasicList);
