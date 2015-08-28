projeto de autorização e autenticação em golang usando o flow Password credentials(OAuth2)

                         REQUEST TOKEN
Client                   /token (POST)                  Server
                       (URL PARAMS)                            
                                                               
  |                grant_type: password                    |   
  |                client_id: {client_id}                  |   
  |                username: {username}                    |   
  |                password: {password}                    |   
  +------------------------------------------------------> |   
  |                                                        |   
  |          access_token:{token}                          |   
  |          refresh_token:{token}                         |   
  |          expires_in: 3600                              |   
  |          token_type:{token_type}(normally "Bearer")    |   
  | <------------------------------------------------------+   
  |                                                        |   
                                                               
