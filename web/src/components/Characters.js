import React from 'react';
import Table from "./Table"
import {getAllCharacters} from "../actions/backend/axios";

const cols = [
    {
        "name": "id",
        "visualName": "ID"
    }, {
        "name": "name",
        "visualName": "Character"
    }, {
        "name": "fun_facts",
        "visualName": "Curiosidades"
    }
];

export default () => {
    const [allCharacters, setAllCharacters] = React.useState([{"name": "No Character", "id": 1, "fun_facts": "no fun"}]);

    React.useEffect(() => {
        let characters = getAllCharacters();
        characters.then(chs => {
            setAllCharacters(chs.data.map(ch =>
                ({
                    "id": ch.id,
                    "name": ch.name,
                    "fun_facts": "NA"
                })
            ))
        })
    }, []);

    console.log("ALL CHARACTERS:", allCharacters)

    return (
        <Table cols={cols} rows={allCharacters}/>
    )
}