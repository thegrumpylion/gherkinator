Feature: List created invoices
  As a user i want to see the list of my created invoices

  Scenario: View invoice
    Given I am on the list screen
    When I click on a specific invoice
    Then I go to the specific invoice details screen

  Scenario: Sort invoices by name
    Given I am on the list screen
    When I click on the sort-by drop-down
    And select by-name
    Then the list must be sorted by name ascending

  Scenario: Sort invoices by date created
    Given I am on the list screen
    When I click on the sort-by drop-down
    And select by-date-created
    Then the list must be sorted by date created descending

  Scenario: Download invoice
    Given I am on the list screen
    When I click on the download icon of a specific invoice
    Then the specific invoice must be downloaded as PDF

  Scenario: Delete invoice
    Given I am on the list screen
    When I click on the delete icon of a specific invoice
    Then the specific invoice must be deleted

  Scenario: Create new invoice
    Given I am on the list screen
    When I click on the create-new icon
    Then I go to the create invoice screen

  Scenario: Edit invoice
    Given I am on the list screen
    When I click the edit button of a specific invoice
    Then I go to the specific invoice edit screen
