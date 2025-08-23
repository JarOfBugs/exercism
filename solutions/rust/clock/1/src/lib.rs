use core::fmt;

#[derive(PartialEq, Debug)]
pub struct Clock {
    hours: i32,
    minutes: i32,
}

impl Clock {
    pub fn new(hours: i32, minutes: i32) -> Self {
        let mut c = Clock {
            hours: hours,
            minutes: minutes,
        };
        c.fix_time();
        c
    }

    pub fn add_minutes(mut self, minutes: i32) -> Self {
        self.minutes += minutes;
        self.fix_time();
        self
    }

    fn fix_time(&mut self) {
        if self.minutes >= 60 || self.minutes < 0 {
            self.hours += self.minutes.div_euclid(60);
            self.minutes = self.minutes.rem_euclid(60);
        }
        if self.hours >= 24 || self.hours < 0 {
            self.hours = self.hours.rem_euclid(24);
        }
    }
}

impl fmt::Display for Clock {
    fn fmt(&self, f: &mut fmt::Formatter<'_>) -> fmt::Result {
        write!(f, "{:0>2}:{:0>2}", self.hours, self.minutes)
    }
}
