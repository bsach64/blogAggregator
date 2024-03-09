Supports the following endpoints:
1.  `GET /v1/readiness`
    Returns a status OK
    ```
    {
      "status": "ok"
    }
    ```
2. `POST /v1/users` with the following request body
    ```
    {
      "name": "Bhavik"
    }
    ```
    Creates a new user and returns:
   ```
   {
      "id": "3f8805e3-634c-49dd-a347-ab36479f3f83",
      "created_at": "2023-09-01T00:00:00Z",
      "updated_at": "2023-09-01T00:00:00Z",
      "name": "Bhavik"
      "api_key": "cca9688383ceaa25bd605575ac9700da94422aa397ef87e765c8df4438bc9942"
   }
   ```
3. `GET /v1/users` which returns user info when provided with a `Authorization: ApiKey <key>` header.
4. `POST /v1/feeds` (requires auth) with the request body
    ```
    {
      "name": "NAME OF WEBPAGE",
      "url": "URL of website's RSS feed"
    }
    ```
    Creates a new feed and returns:
   ```
   {
      "id": "3f8805e3-634c-49dd-a347-ab36479f3f83",
      "created_at": "2023-09-01T00:00:00Z",
      "updated_at": "2023-09-01T00:00:00Z",
      "name": "NAME OF WEBPAGE",
      "url": "URL",
      "user_id": "UUID"
   }
   ```
5. `POST /v1/feed_follows` (requires auth) with the request body
  ```
  {
    "feed_id": "4a82b372-b0e2-45e3-956a-b9b83358f86b"
  }
  ```
  Create a feed follow for the user and returns
  ```
  {
    "id": "c834c69e-ee26-4c63-a677-a977432f9cfa",
    "feed_id": "4a82b372-b0e2-45e3-956a-b9b83358f86b",
    "user_id": "0e4fecc6-1354-47b8-8336-2077b307b20e",
    "created_at": "2017-01-01T00:00:00Z",
    "updated_at": "2017-01-01T00:00:00Z"
  }
  ```
Also, has a background scraper that fetches any new feeds from the blogs in the database, periodically.
  
