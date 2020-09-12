import React from 'react';
import {Button, Form, FormGroup} from 'reactstrap';
import ProfileInput from "../shared/profile-input";
import {Link} from "react-router-dom";

const RegisterNewProfile = (props) => {
    return <div>
        <ProfileInput {...props} isRegisteringNew={true}/>
        <Form className={'col-lg-6 offset-lg-3'}>
            <FormGroup>
                <Button color="primary">Submit</Button>
                <Link to={'/login'}>Login</Link>
            </FormGroup>
        </Form>
    </div>
}

export default RegisterNewProfile;