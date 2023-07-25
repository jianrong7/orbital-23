import React from "react";
import Home from "./page";

describe("<Home />", () => {
  it("renders", () => {
    // see: https://on.cypress.io/mounting-react
    cy.mount(<Home />);
    cy.get(".font-semibold").contains("Postman Emulator");
  });

  it("invalid url and valid json", () => {
    cy.mount(<Home />);
    cy.get("#url").clear();
    cy.get("#url").type("INVALID URL");
    cy.get("p").clear();
    cy.get("p").type(
      JSON.stringify({
        name: "morpheus",
        job: "leader",
      }),
      { parseSpecialCharSequences: false }
    );
    cy.get(".gap-4 > .bg-blue-500").click();
    cy.get(".flex-col.gap-4 > :nth-child(3)", { timeout: 1000 }).should(
      "not.exist"
    );
  });

  it("valid url and invalid json", () => {
    cy.mount(<Home />);
    cy.get("#url").clear();
    cy.get("#url").type("https://reqres.in/api/users");
    cy.get("p").clear();
    cy.get("p").type("INVALID JSON");
    cy.get(".gap-4 > .bg-blue-500").click();
    cy.get(".flex-col.gap-4 > :nth-child(3)", { timeout: 1000 }).should(
      "not.exist"
    );
  });

  it("invalid url and invalid json", () => {
    cy.mount(<Home />);
    cy.get("#url").clear();
    cy.get("#url").type("https://reqres.in/api/users");
    cy.get("p").clear();
    cy.get("p").type("INVALID JSON");
    cy.get(".gap-4 > .bg-blue-500").click();
    cy.get(".flex-col.gap-4 > :nth-child(3)", { timeout: 1000 }).should(
      "not.exist"
    );
  });

  it("valid url and valid json", () => {
    cy.mount(<Home />);
    cy.get("#url").clear();
    cy.get("#url").type("https://reqres.in/api/users");
    cy.get("p").type(
      JSON.stringify({
        name: "morpheus",
        job: "leader",
      }),
      { parseSpecialCharSequences: false }
    );
    cy.get(".gap-4 > .bg-blue-500").click();
    cy.get(".flex-col.gap-4 > :nth-child(3)", { timeout: 1500 })
      .should("be.visible")
      .and("contain", "Response");
    cy.get(":nth-child(3) > .flex > :nth-child(2)").contains("Status: 2");
  });
});
