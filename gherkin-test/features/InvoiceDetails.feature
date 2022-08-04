Feature: Invoice details

  Scenario: Go back
    Given I am on the invoice details screen
    When I click the back button
    Then I go to the invoice list screen

  Scenario: Edit invoice
    Given I am on the invoice details screen
    When I click the edit button
    Then I go to the edit invoice screen

  Scenario: Delete invoice
    Given I am on the invoice details screen
    When I click the delete button
    Then the invoice must be deleted
    And I go to the invoice list screen

  Scenario: Download invoice
    Given I am on the invoice details screen
    When I click on the download icon
    Then invoice must be downloaded as PDF
