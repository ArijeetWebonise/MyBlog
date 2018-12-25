import React, { Component } from 'react';
import { connect } from 'react-redux';
import _ from 'lodash';
import PropType from 'prop-types';
import 'bootstrap/dist/css/bootstrap.min.css';
import { Login } from '../../action/login';
import './Login.sass';

class LoginComponent extends Component {
  constructor(props) {
    super(props);
    this.state = {
      email: '',
      password: '',
      error: [],
    };

    this.OnChangeEmail = this.OnChangeEmail.bind(this);
    this.OnChangePassword = this.OnChangePassword.bind(this);
    this.onSubmit = this.onSubmit.bind(this);
  }

  componentDidMount() {
    const bodyTags = document.getElementsByTagName('body');
    _.each(bodyTags, bodyTag => bodyTag.setAttribute('id', 'loginPage'));
  }

  componentWillUnmount() {
    const bodyTags = document.getElementsByTagName('body');
    _.each(bodyTags, bodyTag => bodyTag.removeAttribute('id'));
  }

  onSubmit(event) {
    const isValid = this.OnValidate();
    event.preventDefault();
    if (!isValid) {
      return false;
    }
    const { email, password } = this.state;

    this.props.Login({ email, password });

    return true;
  }

  OnValidate() {
    const { state } = this;
    const error = [];
    if (_.isEmpty(state.email)) {
      error.push('Email can\'t be empty');
    }

    if (_.isEmpty(state.password)) {
      error.push('Password can\'t be empty');
    }

    if (!_.isEmpty(error)) {
      this.setState({ error });
      return false;
    }

    return true;
  }

  OnChangeEmail(email) {
    this.setState({ email });
  }

  OnChangePassword(password) {
    this.setState({ password });
  }

  render() {
    const { error } = this.state;

    return (
      <div className="container login-container">
        <div className="login-form align-middle">
          <div className="main-div">
            <div className="panel">
              <h2>Admin Login</h2>
              <p>Please enter your email and password</p>
            </div>
            {
              !_.isEmpty(error) && (
                <ul className="error well">
                  {
                    _.map(error, (err, key) => <li key={key}>{err}</li>)
                  }
                </ul>
              )
            }
            <form
              id="Login"
              onSubmit={this.onSubmit}
              >
              <div className="form-group">
                <input
                  type="email"
                  className="form-control"
                  id="inputEmail"
                  placeholder="Email Address"
                  // value={state.email}
                  onChange={(e) => { this.OnChangeEmail(e.currentTarget.value); }}
                  />
              </div>
              <div className="form-group">
                <input
                  type="password"
                  className="form-control"
                  id="inputPassword"
                  placeholder="Password"
                  // value={state.password}
                  onChange={(e) => { this.OnChangePassword(e.currentTarget.value); }}
                  />
              </div>
              <div className="forgot">
                <a href="reset.html">Forgot password?</a>
              </div>
              <button
                type="submit"
                className="btn btn-primary"
                >
                Login
              </button>
            </form>
          </div>
        </div>
      </div>
    );
  }
}

LoginComponent.propTypes = {
  Login: PropType.func.isRequired,
};

const mapStateToProps = state => ({
  login: state.login,
});

const mapActionToProps = dispatch => ({
  Login: ({ user, pass }) => { dispatch(Login(dispatch, { user, pass })); },
});

export default connect(mapStateToProps, mapActionToProps)(LoginComponent);
