# driveorfly
I came up with this idea as i flew from germany to spain and wondered if it'd be better to drive.
So this repository tries to make it possible for people to check wether they should drive or fly to their destination.

To be able to decide which would be better, in *Version 1.0* the user must provide some informations:
- The origin city and coutry of the user and the destination city and country - it is also possible to only provide country, but the outcome maybe significantly different than with cities. 

# Functionality
it calculates for both options the co2 emissions based on co2 emission per 100km and driven/flight-km. Then it tries to suggest the better solution, if both are equally, it tries to determine the time consumption and price of both and decide on one option.

# Privacy
currently no data is stored. In *version three* however there will be an statistic section and for this the saved co2 emissions are stored(without any userdata).

# Disclaimer
This Repo is currently under construction

# Install
```
git clone https://www.github.com/NXZAS8CA/driveorfly.git
cd driveorfly
go run driveorfly
```

# Contribution
Feel free to open an PR, for major changes please open an issue first so we can discuss the changes.
