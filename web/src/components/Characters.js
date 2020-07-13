import React from 'react';
import {Button, Collapse} from "reactstrap";
import Table from "./Table"
import {getAllCharacters} from "../actions/backend/axios";
import Phrases from "./Phrases";

const cols = [
    {
        "name": "id",
        "visualName": "#"
    },
    {
        "name": "content",
        "visualName": "Character"
    }
];

export default () => {
    const [allCharacters, setAllCharacters] = React.useState([]);
    const [rows, setRows] = React.useState([])
    const [open, setOpen] = React.useState({});

    const toggleOpen = (characterId) => {
        setOpen({
            ...open,
            [characterId]: !open[characterId]
        });
    }

    React.useEffect(() => {
        let characters = getAllCharacters();
        characters.then(chs => {
            setAllCharacters(chs.data)
        }).catch(err => {
            console.log("ERROR: ", err)
        })
    }, [])

    React.useEffect(() => {
        setRows(allCharacters.map(ch => {
                let rowContent =
                    <div>
                        <Button color="link" onClick={() => toggleOpen(ch.id)}>
                            {ch.name}
                        </Button>
                        <Collapse isOpen={open[ch.id]}>
                            <Phrases characterId={ch.id}/>
                        </Collapse>
                    </div>;

                return {
                    "id": ch.id,
                    "content": rowContent
                }
            }
        ))
    }, [allCharacters, open]);

    return (
        <div className="container-sm" align="left">
            <Table cols={cols} rows={rows}/>
        </div>
    )
}