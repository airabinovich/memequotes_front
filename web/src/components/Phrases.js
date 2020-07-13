import React from 'react';
import {ListGroup, ListGroupItem} from 'reactstrap';
import {getAllPhrasesForCharacter} from "../actions/backend/axios";

export default (props) => {
    const [phrases, setPhrases] = React.useState([])

    React.useEffect(() => {
        let phrasesPromise = getAllPhrasesForCharacter(props.characterId);

        phrasesPromise.then(phs => {
            setPhrases(phs.data.map(ph => {
                    return <ListGroupItem>{ph.content}</ListGroupItem>
                }
            ))
        }).catch(err => {
            console.log("ERROR: ", err)
        })
    }, []);

    return (
        <ListGroup>
            {phrases}
        </ListGroup>
    )
}