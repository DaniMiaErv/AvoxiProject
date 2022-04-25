# AvoxiProject
An api microservice to handle request.

Scenario (note, this is fictional, but gives a sense of the types of requests we might encounter):
Our product team has heard from several customers that we restrict users to logging in to their UI accounts from selected countries to prevent them from outsourcing their work to others. For an initial phase one we're not going to worry about VPN connectivity, only the presented IP address.
The team has designed a solution where the customer database will hold the white listed countries and the API gateway will capture the requesting IP address, check the target customer for restrictions, and send the data elements to a new service you are going to create. 

************The new service will be an HTTP-based API that receives an IP address and a white list of countries. The API should return an indicator if the IP address is within the listed countries. You can get a data set of IP addresses to country mappings from https://dev.maxmind.com/geoip/geoip2/geolite2/.

