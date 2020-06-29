import React, {useState} from 'react';
import {Button, Form, FormGroup, Label, Input, FormText} from 'reactstrap';
import ButtonGroup from "reactstrap/es/ButtonGroup";
import ProfileInput from "../shared/profile-input";
import {Link} from "react-router-dom";

const ProfileEdit = (props) => {
    return <div>
        <ProfileInput {...props} isRegisteringNew={false}/>
        <Form className={'col-lg-6 offset-lg-3'}>
            <FormGroup>
                <ButtonGroup>
                    <Button color="primary">Submit</Button>
                    <Link to={'/'}>Cancel</Link>
                </ButtonGroup>
            </FormGroup>
        </Form>
    </div>
}

export default ProfileEdit;