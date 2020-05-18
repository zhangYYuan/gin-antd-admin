import React, { useRef, useState, useEffect } from 'react';
import { connect } from 'umi';
import { Card, Col, Row, Tree, Alert } from 'antd';
import { GridContent } from '@ant-design/pro-layout';
import CreateDepart from './components/CreateDepart';

export const Depart = props => {
  const [current, setCurrent] = useState('已选: ')

  const {
    dispatch,
    sysDepart: { departList },
  } = props;

  useEffect( () => {
    dispatch({
      type: 'sysDepart/fetch'
    });
  }, [])

  const onSelect = (selectedKeys, info) => {
    setCurrent(`已选: ${info.node.title}`)
  };

  return (
    <GridContent>
      <Row gutter={12}>
        <Col style={{ marginBottom: 24 }} lg={7} md={24} xs={24}>
          <Card
            bordered={false}
          >
            <div style={{ marginBottom: 24 }}>
              <Alert type="info" showIcon message={current} />
            </div>
            <Tree
              onSelect={onSelect}
              treeData={departList}
            />
          </Card>
        </Col>
        <Col lg={10} md={24} xs={24}>
          <Card
            bordered={false}
          >
            <CreateDepart />
          </Card>
        </Col>
      </Row>
    </GridContent>
  )
}


export default connect(({ sysDepart }) => ({
  sysDepart
}))(Depart);
