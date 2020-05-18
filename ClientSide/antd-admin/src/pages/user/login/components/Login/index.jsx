import { Tabs, Form } from 'antd';
import React, { useState } from 'react';
import classNames from 'classnames';
import LoginContext from './LoginContext';
import LoginItem from './LoginItem';
import LoginSubmit from './LoginSubmit';
import styles from './index.less';

const Login = props => {
  const { className } = props;
  return (
    <LoginContext.Provider>
      <div className={classNames(className, styles.login)}>
        <Form
          form={props.from}
          onFinish={values => {
            if (props.onSubmit) {
              props.onSubmit(values);
            }
          }}
        >
          { props.children }
        </Form>
      </div>
    </LoginContext.Provider>
  );
};

Login.Submit = LoginSubmit;
Login.UserName = LoginItem.UserName;
Login.Password = LoginItem.Password;

export default Login;
