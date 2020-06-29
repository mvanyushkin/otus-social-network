import React from 'react';
import {Button, Form, FormGroup, Label, Input, Col, Jumbotron, Container} from 'reactstrap';
import {Link} from "react-router-dom";

const ProfileView = (props) => {
    return <div className={'col-lg-6 offset-lg-3'}>

        <Jumbotron fluid>
            <Container fluid>
                <h1 className="display-3">{props.firstName} {props.lastName}</h1>
                <p className="lead">This is your own page!</p>
            </Container>
        </Jumbotron>

        <Form>

            <FormGroup row>
                <Label sm={2} for="loginEmail">Email</Label>
                <Col sm={10}>
                    <Input type="email" name="loginEmail" id="loginEmail"
                           value={props.loginEmail} disabled/>
                </Col>
            </FormGroup>

            <FormGroup row>
                <Label sm={2} for="firstName">First Name</Label>
                <Col sm={10}>
                    <Input type="password" name="password" id="firstName"
                           value={props.firstName}
                           disabled/>
                </Col>
            </FormGroup>

            <FormGroup row>
                <Label sm={2} for="lastName">Last Name</Label>
                <Col sm={10}>
                    <Input type="password" name="password" id="lastName" value={props.lastName}
                           disabled/>
                </Col>
            </FormGroup>

            <FormGroup row>
                <Label sm={2} for="HobbyDescription">Hobby Description</Label>
                <Col sm={10}>
                    <Input type="textarea" name="hobbyDescription" id="hobbyDescription" value={props.hobbyDescription}
                           disabled/>
                </Col>
            </FormGroup>

            <FormGroup row>
                <Label sm={2} for="gender">Gender</Label>
                <Col sm={10}>
                    <Input type="select" name="gender" id="gender" value={props.gender} disabled>
                        <option>Male</option>
                        <option>Female</option>
                        <option>Prefer to keep in a secret</option>
                    </Input>
                </Col>
            </FormGroup>

            <FormGroup row>
                <Label sm={2} for="city">City</Label>
                <Col sm={10}>
                    <Input type="text" name="city" id="city" value={props.city}
                           disabled/>
                </Col>
            </FormGroup>

            <FormGroup row>
                <Label sm={2} for="age">Age</Label>
                <Col sm={10}>
                    <Input type="number" name="age" id="age" value={props.age} disabled/>
                </Col>
            </FormGroup>

            <FormGroup>
                <Button color={'success'}>
                    <Link to="/edit">Edit</Link>
                </Button>
            </FormGroup>
        </Form>
    </div>
}

export default ProfileView;