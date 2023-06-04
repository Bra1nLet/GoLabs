import React from 'react';
import { useRef, useState } from 'react';
import {Authorize}  from "../../hooks";
import {ExitFromAccount, GetToken} from "../../hooks/Auth";
import './Drive.css'

const RenameDrive = () => {
    const path = useRef("");
    const nameFile = useRef("");
    const newNameFile = useRef("");

    const handleButtonRename = () => {
        const formData = new FormData();
        formData.append('path', path.current.value);
        formData.append('filename', nameFile.current.value);
        formData.append('newname', newNameFile.current.value);
        const authOptions = {
            method: "POST",
            body: formData,
            headers: {
                'Authorization': GetToken()
            },
        };

        fetch('/drive/rename', authOptions);
    };


    return (
        <div>
            <h1>Rename</h1>
            <label htmlFor="path">Path</label>
            <br />
            <input ref={path} name="path" type="text" />
            <br />
            <label htmlFor="name">Name</label>
            <br />
            <input ref={nameFile} name="name" type="text" />
            <br />
            <label htmlFor="newname">New Name</label>
            <br />
            <input ref={newNameFile} name="newname" type="text" />
            <br />
            <button onClick={handleButtonRename}>Rename</button>
        </div>
    );
};
export default RenameDrive;