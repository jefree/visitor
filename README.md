# VISITOR
*Test your web app with hundreds of requests at the same time*

## Main Idea

Automating stress navigation tests using ruby as scripting language and a powerful API build over Selenium by runnig hundreds of agents at the same time.

## Example (not available yet **)
** API can change at any time

```ruby
def start_new_session(visitor)
    visitor.visit('http://myapp.com/')
    
    visitor.click_on 'Log In'
    
    visitor.fill_in 'username', with: 'myusername'
    visitor.fill_in 'password', with: '******'
    
    visitor.click_on 'Send'
end

Visitor.parallelize(start_new_session, 1000) # runs 1000 concurrent logins
```
