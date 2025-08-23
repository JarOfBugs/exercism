// The code below is a stub. Just enough to satisfy the compiler.
// In order to pass the tests you can add-to or change any of this code.

const EARTH_YEAR_IN_SECONDS: f64 = 31_557_600.0;

#[derive(Debug)]
pub struct Duration {
    seconds: f64,
}

impl From<u64> for Duration {
    fn from(s: u64) -> Self {
        Duration { seconds: s as f64 }
    }
}

pub trait Planet {
    fn years_during(d: &Duration) -> f64 {
        d.seconds / (EARTH_YEAR_IN_SECONDS * Self::get_orbit_multiplier())
    }

    fn get_orbit_multiplier() -> f64;
}

macro_rules! create_planet {
    ($name:ident, $orbit:literal) => {
        pub struct $name;
        impl Planet for $name {
            fn get_orbit_multiplier() -> f64 {
                $orbit
            }
        }
    };
}

create_planet!(Earth, 1.0);
create_planet!(Mercury, 0.2408467);
create_planet!(Venus, 0.61519726);
create_planet!(Mars, 1.8808158);
create_planet!(Jupiter, 11.862615);
create_planet!(Saturn, 29.447498);
create_planet!(Uranus, 84.016846);
create_planet!(Neptune, 164.79132);
