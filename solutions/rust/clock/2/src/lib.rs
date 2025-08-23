use core::fmt;

#[derive(PartialEq, Debug)]
pub struct Clock {
    hours: i32,
    minutes: i32,
}

impl Clock {
    pub fn new(mut hours: i32, mut minutes: i32) -> Self {
        if minutes >= 60 || minutes < 0 {
            hours += minutes.div_euclid(60);
            minutes = minutes.rem_euclid(60);
        }
        if hours >= 24 || hours < 0 {
            hours = hours.rem_euclid(24);
        }
        Clock { hours, minutes }
    }

    pub fn add_minutes(self, minutes: i32) -> Self {
        Self::new(self.hours, self.minutes + minutes)
    }
}

impl fmt::Display for Clock {
    fn fmt(&self, f: &mut fmt::Formatter<'_>) -> fmt::Result {
        write!(f, "{:0>2}:{:0>2}", self.hours, self.minutes)
    }
}
