use std::cmp::Ordering;

const INVALID_POSITION: i32 = -1;

#[derive(Debug)]
pub struct Token {
    pub character: char,
    pub first_position: i32,
    pub count: u32,
}

impl Default for Token {
    fn default() -> Token {
        Token {
            character: '?',
            first_position: -1,
            count: 0,
        }
    }
}

impl Token {
    fn cmp(&self, rhs: &Token) -> Ordering {
        if self.first_position == INVALID_POSITION && rhs.first_position == INVALID_POSITION {
            return Ordering::Equal;
        }
        if self.first_position == INVALID_POSITION && rhs.first_position != INVALID_POSITION {
            return Ordering::Greater;
        }
        if self.first_position != INVALID_POSITION && rhs.first_position == INVALID_POSITION {
            return Ordering::Less;
        }

        if self.count == rhs.count {
            // if tied, check lexicographic ordering
            return self.character.cmp(&rhs.character);
        }

        // if self.first_position > rhs.first_position {
        //     return Ordering::Greater;
        // }
        // if self.first_position < rhs.first_position {
        //     return Ordering::Less;
        // }

        self.count.cmp(&rhs.count).reverse()
    }
}

#[derive(Debug)]
pub struct Room {
    pub name: String,
    pub sector_id: u32,
    pub actual_hash: String,
    pub token_count: [Token; 26],
}

impl Default for Room {
    fn default() -> Room {
        Room {
            name: String::new(),
            sector_id: 0,
            actual_hash: String::new(),
            token_count: Default::default(),
        }
    }
}

impl Room {
    const HASH_LENGTH: usize = 5;

    pub fn parse(&mut self, line: &str) {
        self.name = String::from(line);
        self.parse_hash(line);
        self.parse_sector_id(line);
        self.count_tokens(line);
    }

    fn parse_hash(&mut self, line: &str) {
        // hash is the last 5 chars
        let n = line.len();
        let end = n - 1; // final char == ]
        if &line[end..] != "]" {
            panic!("");
        }
        let start = end - Room::HASH_LENGTH; // start char == ]
        if &line[start - 1..start] != "[" {
            panic!("");
        }
        self.actual_hash = String::new();
        self.actual_hash += &line[start..end];
    }

    fn parse_sector_id(&mut self, line: &str) {
        let end = line.len() - 2 - Room::HASH_LENGTH; // start char == ]
        if &line[end..end + 1] != "[" {
            panic!("");
        }
        let mut start = end;
        for i in 0..line.len() {
            if &line[start - i..start - i + 1] == "-" {
                start = start - i + 1;
                break;
            }
        }
        if start == end {
            panic!("");
        }
        self.sector_id = line[start..end].parse().expect("Not a valid sector id");
    }

    fn count_tokens(&mut self, line: &str) {
        // reset tokens
        for token in self.token_count.iter_mut() {
            token.count = 0;
        }

        for (i, c) in line.chars().enumerate() {
            if !c.is_alphanumeric() {
                continue;
            }
            if c.is_digit(10) {
                break; // sector id denotes the end of the room id
            }

            let idx = c as usize - 'a' as usize;
            self.token_count[idx].character = c;

            //print!("[{}, {}]: {:?}", i, c, self.token_count[idx]);
            if self.token_count[idx].first_position == INVALID_POSITION {
                self.token_count[idx].first_position = i as i32;
            }
            self.token_count[idx].count += 1;
            //print!("-> {:?}\n", self.token_count[idx]);
        }

        // now sort token_count
        self.token_count.sort_by(|a, b| a.cmp(&b));
        //println!("{:#?}", *self);
    }

    pub fn is_real(&self) -> bool {
        self.actual_hash == self.get_hash()
    }

    pub fn get_hash(&self) -> String {
        let mut ret = String::new();

        for i in 0..Room::HASH_LENGTH {
            ret.push(self.token_count[i].character);
        }

        ret
    }

    pub fn decode(&self) -> String {
        let mut buffer = self.name.clone();

        const NUMBER_OF_LETTERS: u32 = 26;
        const OFFSET: u32 = 'a' as u32;

        let n_iterations = self.sector_id;

        buffer = buffer
            .chars()
            .map(|c| match c {
                '-' => ' ',
                x if x.is_alphabetic() => {
                    ((OFFSET + ((x as u32 - OFFSET) + n_iterations) % NUMBER_OF_LETTERS) as u8)
                        as char
                }
                _ => c,
            })
            .collect();

        buffer
    }
}
