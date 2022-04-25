# AvoxiProject
An api microservice to check if a given IP address comes from a country that is whitelisted.

@param ipAd String 
        a given IP address to check
@param countryList []String 
        a list of countries on a Whitelist
@return passFail boolean 
        returns true if the IP address comes from a country in the list
        returns false if the IP address comes from a country outside of the list
@return err Error
        returns an error
