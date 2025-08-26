pub fn recite(start_bottles: u32, take_down: u32) -> String {
    let numbers = [
        "no", "One", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten",
    ];
    let mut result: Vec<String> = vec![];
    let mut index = start_bottles;
    let get_bottle_word = |num| -> &str {
        match num {
            1 => "bottle",
            _ => "bottles",
        }
    };
    for _ in 0..take_down {
        for _ in 0..2 {
            result.push(format!(
                "{} green {} hanging on the wall,",
                numbers[index as usize],
                get_bottle_word(index)
            ));
        }
        result.push("And if one green bottle should accidentally fall,".to_string());
        index -= 1;
        result.push(format!(
            "There'll be {} green {} hanging on the wall.\n",
            numbers[index as usize].to_lowercase(),
            get_bottle_word(index)
        ));
    }

    result.join("\n")
}
