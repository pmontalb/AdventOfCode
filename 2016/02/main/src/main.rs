use std::fs::File;
use std::io::Read;
use std::io::Write;

#[derive(Debug)]
enum Move {
    Up,
    Down,
    Left,
    Right,
}
impl Move {
    fn parse(token: &char) -> Move {
        match token {
            'U' => Move::Up,
            'D' => Move::Down,
            'L' => Move::Left,
            'R' => Move::Right,
            _ => unimplemented!(),
        }
    }
}

const MIN_POSITION: i32 = -1;
const MAX_POSITION: i32 = 1;
#[derive(Debug)]
struct KeyPadButton {
    x: i32,
    y: i32,
}

impl KeyPadButton {
    fn _is_valid(x: i32, y: i32) -> bool {
        x >= MIN_POSITION && x <= MAX_POSITION && y >= MIN_POSITION && y <= MAX_POSITION
    }
    fn follow(&mut self, instruction: Move) {
        let mut new_x = self.x;
        let mut new_y = self.y;
        match instruction {
            Move::Up => new_y -= 1,
            Move::Down => new_y += 1,
            Move::Left => new_x -= 1,
            Move::Right => new_x += 1,
        }
        if KeyPadButton::_is_valid(new_x, new_y) {
            self.x = new_x;
            self.y = new_y;
        }
    }

    fn get_id(&self) -> i32 {
        match self.y {
            -1 => match self.x {
                -1 => 1,
                0 => 2,
                1 => 3,
                _ => unimplemented!(),
            },

            0 => match self.x {
                -1 => 4,
                0 => 5,
                1 => 6,
                _ => unimplemented!(),
            },

            1 => match self.x {
                -1 => 7,
                0 => 8,
                1 => 9,
                _ => unimplemented!(),
            },

            _ => unimplemented!(),
        }
    }
}

#[derive(Debug)]
struct ComplexKeyPadButton {
    x: i32,
    y: i32,
}
impl ComplexKeyPadButton {
    fn _is_valid(x: i32, y: i32) -> bool {
        match y {
            -2 => x == 0,
            -1 => x >= -1 && x <= 1,
            0 => x >= -2 && x <= 2,
            1 => x >= -1 && x <= 1,
            2 => x == 0,
            _ => false,
        }
    }
    fn follow(&mut self, instruction: Move) {
        let mut new_x = self.x;
        let mut new_y = self.y;
        match instruction {
            Move::Up => new_y -= 1,
            Move::Down => new_y += 1,
            Move::Left => new_x -= 1,
            Move::Right => new_x += 1,
        }
        if ComplexKeyPadButton::_is_valid(new_x, new_y) {
            self.x = new_x;
            self.y = new_y;
        }
    }

    fn get_id(&self) -> i32 {
        match self.y {
            -2 => match self.x {
                0 => 0x1,
                _ => unimplemented!(),
            },

            -1 => match self.x {
                -1 => 2,
                0 => 3,
                1 => 4,
                _ => unimplemented!(),
            },

            0 => match self.x {
                -2 => 5,
                -1 => 6,
                0 => 7,
                1 => 8,
                2 => 9,
                _ => unimplemented!(),
            },

            1 => match self.x {
                -1 => 0xA,
                0 => 0xB,
                1 => 0xC,
                _ => unimplemented!(),
            },

            2 => match self.x {
                0 => 0xD,
                _ => unimplemented!(),
            },

            _ => unimplemented!(),
        }
    }
}

fn main() {
    let mut input_file = File::open("../input").expect("Unable to open");
    let mut contents = String::new();
    input_file
        .read_to_string(&mut contents)
        .expect("Unable to read");

    let mut button = KeyPadButton { x: 0, y: 0 };
    let lines = contents.split("\n");

    let mut code = String::new();
    for line in lines {
        for instruction in line.chars() {
            let move_instruction = Move::parse(&instruction);

            //print!("{:?} => ", move_instruction);
            button.follow(move_instruction);
            //println!("{:?}", button);
        }
        code += &button.get_id().to_string();
        //println!("*CODE: {}", code);
    }
    println!("{}", code);
    let mut output1 = File::create(&"output1").expect("Unable to create");
    output1.write_all(code.as_bytes()).expect("Unable to write");

    // now do the same but with the complex keypad
    let mut complex_button = ComplexKeyPadButton { x: -2, y: 0 };
    let lines = contents.split("\n");

    let mut complex_code = String::new();
    for line in lines {
        for instruction in line.chars() {
            let move_instruction = Move::parse(&instruction);

            //print!("{:?} => ", move_instruction);
            complex_button.follow(move_instruction);
            //println!("{:?}", complex_button);
        }
        complex_code += &format!("{:X}", complex_button.get_id());
        //println!("*CODE: {}", complex_code);
    }
    println!("{}", complex_code);
    let mut output2 = File::create(&"output2").expect("Unable to create");
    output2
        .write_all(complex_code.as_bytes())
        .expect("Unable to write");
}
