import React from 'react';


import Login from "./login/login";
import RegisterNewProfile from "./register/register-new-profile";
import ProfileView from "./profile/profile-view";
import {Route, withRouter} from "react-router-dom";
import ProfileEdit from "./profile/profile-edit";

class App extends React.Component {

    constructor(props) {
        super(props);

        this.state = {
            isProfileLoaded: false,
            isAuthenticated: false
        }
    }


    componentDidMount() {
        if (!this.state.isAuthenticated && this.props.history.location.pathname !== '/login') {
            //this.props.history.push("/login");
        }
    }

    render() {


        return (
            <div className="App">
                <Route exact path='/'>
                    <ProfileView />
                </Route>
                <Route exact path='/edit'>
                    <ProfileEdit />
                </Route>
                <Route exact path='/login'>
                    <Login/>
                </Route>
                <Route exact path='/register'>
                    <RegisterNewProfile/>
                </Route>
            </div>
        );
    }
}

export default withRouter(App);
