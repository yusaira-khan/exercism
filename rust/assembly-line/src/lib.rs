// This stub file contains items that aren't used yet; feel free to remove this module attribute
// to enable stricter warnings.
#![allow(unused)]

fn error_rate(speed: u8) -> f64{
    match speed {
        1..=4 => 1.0,
        5..=8 => 0.9,
        9 | 10 => 0.77,
        _ => 0.0
    }
}
pub fn production_rate_per_hour(speed: u8) -> f64 {
    let base_rate :f64 = 221.0;
    error_rate(speed) * base_rate * (speed as f64)
}

pub fn working_items_per_minute(speed: u8) -> u32 {
    (production_rate_per_hour(speed)/60.0)as u32
}
