import React, { useRef, useState, useEffect } from 'react';
import { SetToken } from "../../hooks/Auth";

const Registration = () => {
    const name = useRef(null);
    const username = useRef(null);
    const password = useRef(null);


    const handleButtonClick = () => {
        const authOptions = {
            method: 'POST',
            headers: {
                Accept: 'application/json',
            },
            body: JSON.stringify({
                name: name.current.value,
                username: username.current.value,
                password: password.current.value,
            }),
        };

        fetch('/auth/new-account', authOptions)
            .then(response => response.json())
            .then(response => createAccount())
            .catch(error => {
                console.error('Error:', error);
            });
    };

    function createAccount() {
        window.location.reload(false);
    }

    return (
        <div>
            <h1>Registration</h1>
            <label htmlFor="name">Name</label>
            <br />
            <input ref={name} name="name" type="text" />
            <br />
            <label htmlFor="user">Username</label>
            <br />
            <input ref={username} name="user" type="text" />
            <br />
            <label htmlFor="pass">Password</label>
            <br />
            <input ref={password} name="pass" type="password" />
            <br />
            <button onClick={handleButtonClick}>Create new account</button>
        </div>
    );
};

export default Registration;
