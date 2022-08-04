Feature: Create invoice

  Scenario: Go back
    Given I am on the create invoice screen
    When I click the back button
    Then I go to the invoice list screen

  Scenario: Create correct invoice  
    Given I am on the create invoice screen
    When I enter everything correct
    And I click the save button
    Then I get a success notification
    And I go to the invoice list screen

  Scenario: Create invoice with missing fields
    Given I am on the create invoice screen
    And I partially complete the invoice
    When I click the save button
    Then I see the missing fields errors

  Scenario: Create invoice with existing id
    Given I am on the create invoice screen
    And I enter an existing id in the id field
    When I click the save button
    Then I see the error invoice with id already exists

  Scenario: Create invoice with no line items
    Given I am on the create invoice screen
    And I enter everything corect but no line items
    When I click the save button
    Then I see the error line item missing

  Scenario: Create invoice with custom date
    Given I am on the create invoice screen
    And I enter everything corect
    And I change the date
    When I click the save button
    Then I get a success notification
    And I go to the invoice list screen

  Scenario: Add new line item
    Given I am on the create invoice screen
    And I enter everything corect
    When I click the add new line item button
    Then I can add add a new line item

  Scenario: Remove line item
    Given I am on the create invoice screen
    And I enter everything corect
    When I click the delete line item button
    Then I the line item must be removed from the list
