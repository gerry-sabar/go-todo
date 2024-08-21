*** Settings ***
Library    RequestsLibrary

*** Variables ***
${BASE_URL}    http://localhost:8080

*** Test Cases ***
Get Todo
    [Documentation]    A test case to get a single todo
    Create Session    jsonplaceholder    ${BASE_URL}    verify=false
    ${response}=    GET On Session    jsonplaceholder    v1/todos/1
    Should Be Equal As Numbers    ${response.status_code}    200
    ${json}    Set Variable    ${response.json()}
    Log to console    ${json['data']['id']}
    Should Be Equal As Integers  ${json['data']['id']}     1
    Should Be Equal As Strings  ${json['data']['title']}     Testing
    Should Be Equal As Strings  ${json['data']['status']}    Doing


Create Todo
    [Documentation]    A test case to get a single todo
    Create Session    jsonplaceholder    ${BASE_URL}    verify=false
    ${request_body}    Create Dictionary
    ...    title=Testing Input
    ...    status=Doing

    Create Session    jsonplaceholder      ${BASE_URL}    verify=false
    ${response}    POST On Session  jsonplaceholder   v1/todos
    ...    json=${request_body}
    ...    expected_status=200
    ${json}    Set Variable    ${response.json()}
    Log to console    ${json['data']['id']}
