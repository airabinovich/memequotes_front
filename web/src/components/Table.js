import React from 'react';

export default (props) => {
    let columns = props.cols.map(column =>
        <th scope="col" className="table-col table-heading-col">
            <span className="table-content">{column["visualName"]}</span>
        </th>)
    let rows = props.rows.map((row, key) => {
        return (
            <tr key={key}>
                {
                    props.cols.map(column => {
                        return <td>{row[column.name]}</td>
                    })
                }
            </tr>);
    })
    return (
        <table className="table table-striped" width="100%">
            <thead className="thead-light">
            <tr>{columns}</tr>
            </thead>
            <tbody>{rows}</tbody>
        </table>
    );
}

