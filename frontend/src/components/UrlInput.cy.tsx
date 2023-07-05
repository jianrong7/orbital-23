import React from "react";
import UrlInput from "./UrlInput";
import { existsSync } from "fs";

describe("<UrlInput />", () => {
  it("renders", () => {
    // see: https://on.cypress.io/mounting-react
    cy.mount(<UrlInput />);
    cy.get("#url").should("exist");
    cy.get(".bg-blue-500").contains("Send POST Request");
  });
});
