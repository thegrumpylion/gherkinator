Feature: Edit Invoice

  Scenario: Go back
    Given I am on the edit invoice screen
    When I click the back button
    Then I go to the screen I came from

  Scenario: Save edited invoice
    Given I am on the edit invoice screen
    And I change some fields correctly
    When I click the save button
    Then the updated invoice is saved
    And I go to the invoce list screen

  Scenario: Save invoice with missing fields
    Given I am on the edit invoice screen
    And I partially complete the invoice
    When I click the save button
    Then I see the missing fields errors

  Scenario: Save invoice with existing id
    Given I am on the create invoice screen
    And I enter an existing id in the id field
    When I click the save button
    Then I see the error invoice with id already exists

  Scenario: Save invoice with no line items
    Given I am on the create invoice screen
    And I enter everything corect but no line items
    When I click the save button
    Then I see the error line item missing
