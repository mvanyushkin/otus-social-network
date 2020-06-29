import React from 'react';

import { Button, Form, FormGroup, Label, Input } from 'reactstrap';
import {Link} from "react-router-dom";

function Login() {
    return (
        <Form className={'col-lg-6 offset-lg-3'}>
            <FormGroup>
                <Label for="exampleEmail">Email</Label>
                <Input type="email" name="email" id="exampleEmail" placeholder="with a placeholder" />
            </FormGroup>
            <FormGroup>
                <Label for="examplePassword">Password</Label>
                <Input type="password" name="password" id="examplePassword" placeholder="password placeholder" />
            </FormGroup>
            <Button>Submit</Button>
            <Link to={'/register'}>Register a new account</Link>
        </Form>
    );
}

export default Login;
