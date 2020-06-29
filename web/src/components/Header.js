import React from "react";
import {Navbar, NavbarBrand} from "reactstrap";

export default () => {
    return (
        <React.Fragment>
            <Navbar dark color="primary" sticky="top">
                <div className="container">
                    <NavbarBrand href="/">MemeQuotes</NavbarBrand>
                </div>
            </Navbar>
        </React.Fragment>
    );
};