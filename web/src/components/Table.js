import React from 'react';

export default (props) => {
    let columns = props.cols.map(column =>
        <th scope="col" className="table-col table-heading-col" width={`${50. / (props.cols.length)}%`}>
            <span className="table-content">{column["visualName"]}</span>
        </th>)
    let rows = props.rows.map((result, key) => {
        return <tr className="table-row" key={key}> {
            props.cols.map(column => {
                return <td className="table-col">{result[column.name].toString()}</td>
            })
        }
        </tr>;
    })
    return (
        <table className="table table-striped" width="50%">
            <thead className="thead-light">
            <tr>{columns}</tr>
            </thead>
            <tbody className="table-body">{rows}</tbody>
        </table>
    );
}

