import React from 'react';
import {Button, Form, FormGroup, Label, Input, FormText} from 'reactstrap';
import ButtonGroup from "reactstrap/es/ButtonGroup";

const ProfileInput = (props) => {
    return (
        <Form readonly className={'col-lg-6 offset-lg-3'} >
            <FormGroup>
                <Label for="loginEmail">Email</Label>
                <Input type="email" name="loginEmail" id="loginEmail" placeholder="Login email"
                       value={props.loginEmail} disabled={props.disabled}/>
            </FormGroup>
            {
                props.isRegisteringNew &&
                <FormGroup>
                    <Label for="password">Password</Label>
                    <Input type="password" name="password" id="password" placeholder="Password" />
                </FormGroup>
            }

            <FormGroup>
                <Label for="firstName">First Name</Label>
                <Input type="password" name="password" id="firstName" placeholder="First Name" value={props.firstName} disabled={props.disabled}/>
            </FormGroup>
            <FormGroup>
                <Label for="lastName">Last Name</Label>
                <Input type="password" name="password" id="lastName" placeholder="Last Name" value={props.lastName} disabled={props.disabled}/>
            </FormGroup>
            <FormGroup>
                <Label for="HobbyDescription">Hobby Description</Label>
                <Input type="textarea" name="hobbyDescription" id="hobbyDescription" value={props.hobbyDescription} disabled={props.disabled}/>
            </FormGroup>
            <FormGroup>
                <Label for="gender">Gender</Label>
                <Input type="select" name="gender" id="gender" value={props.gender} disabled={props.disabled}>
                    <option>Male</option>
                    <option>Female</option>
                    <option>Prefer to keep in a secret</option>
                </Input>
            </FormGroup>
            <FormGroup>
                <Label for="city">City</Label>
                <Input type="text" name="city" id="city" placeholder="City" value={props.city} disabled={props.disabled}/>
            </FormGroup>

            <FormGroup>
                <Label for="age">Age</Label>
                <Input type="number" name="age" id="age" placeholder="Age" value={props.age} disabled={props.disabled}/>
            </FormGroup>

        </Form>
    );
}

export default ProfileInput;