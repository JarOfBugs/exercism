pub fn annotate(garden: &[&str]) -> Vec<String> {
    if garden.len() == 0 {
        return Vec::new();
    }

    let mut garden_grid: Vec<Vec<u8>> = garden
        .iter()
        .map(|string| string.as_bytes().to_vec())
        .collect();

    for row in 0..garden_grid.len() {
        for col in 0..garden_grid[0].len() {
            if garden_grid[row][col] == b' ' {
                let mut count = 0;
                for diff_row in -1..=1 {
                    for diff_col in -1..=1 {
                        let check_row = (row as isize) + diff_row;
                        let check_col = (col as isize) + diff_col;

                        if check_row >= 0
                            && check_row < (garden_grid.len() as isize)
                            && check_col >= 0
                            && check_col < (garden_grid[0].len() as isize)
                        {
                            if garden_grid[check_row as usize][check_col as usize] == b'*' {
                                count = count + 1;
                                continue;
                            }
                        }
                    }
                }
                if count > 0 {
                    garden_grid[row][col] = b'0' + count;
                }
            }
        }
    }

    garden_grid
        .iter()
        .map(|vec| String::from_utf8(vec.to_owned()).unwrap())
        .collect()
}
