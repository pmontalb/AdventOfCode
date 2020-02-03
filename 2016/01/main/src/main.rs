use std::collections::HashSet;
use std::fs::File;
use std::io::Read;
use std::io::Write;

#[derive(Debug, PartialEq, Eq, Hash, Clone)]
enum Direction {
    None,
    Up,
    Down,
    Left,
    Right,
}
#[derive(Debug, PartialEq, Eq, Hash, Clone)]
struct Coordinate {
    x: i32,
    y: i32,
}
#[derive(Debug, PartialEq, Eq, Hash, Clone)]
struct Position {
    coordinate: Coordinate,
    direction: Direction,
}

#[derive(Debug)]
struct MoveInstruction {
    steps: i32,
    direction: Direction,
}

fn parse(instruction: &mut MoveInstruction, string: &str) {
    match string.as_bytes()[0] {
        b'L' => instruction.direction = Direction::Left,
        b'R' => instruction.direction = Direction::Right,
        _ => unimplemented!(),
    }

    instruction.steps = String::from(&string[1..]).parse().unwrap();
}

fn rotate(origin: &mut Direction, target: &Direction) {
    match target {
        Direction::Left => match origin {
            Direction::Up => *origin = Direction::Left,
            Direction::Left => *origin = Direction::Down,
            Direction::Down => *origin = Direction::Right,
            Direction::Right => *origin = Direction::Up,
            _ => unimplemented!(),
        },
        Direction::Right => match origin {
            Direction::Up => *origin = Direction::Right,
            Direction::Right => *origin = Direction::Down,
            Direction::Down => *origin = Direction::Left,
            Direction::Left => *origin = Direction::Up,
            _ => unimplemented!(),
        },
        _ => unimplemented!(),
    }
}

fn advance(position: &mut Position, steps: i32) {
    match position.direction {
        Direction::Left => position.coordinate.x -= steps,
        Direction::Right => position.coordinate.x += steps,
        Direction::Down => position.coordinate.y -= steps,
        Direction::Up => position.coordinate.y += steps,
        _ => unimplemented!(),
    }
}

fn process(mut position: &mut Position, instruction: &MoveInstruction) {
    rotate(&mut position.direction, &instruction.direction);
    advance(&mut position, instruction.steps);
}

fn get_shortest_path(coordinate: &Coordinate) -> i32 {
    i32::abs(coordinate.x) + i32::abs(coordinate.y)
}

fn main() {
    let mut input_file = File::open("../input").expect("Unable to open");
    let mut contents = String::new();
    input_file
        .read_to_string(&mut contents)
        .expect("Unable to read");

    let instructions = contents.split(", ");

    let mut position = Position {
        coordinate: Coordinate { x: 0, y: 0 },
        direction: Direction::Up,
    };

    let mut visited_places = HashSet::new();
    visited_places.insert(position.coordinate.clone());

    let mut actual_location = position.coordinate.clone();
    let mut has_found_actual_location = false;
    let mut old_position = position.coordinate.clone();

    //println!("Now in {:?}", position);
    for instruction in instructions {
        let mut move_instruction: MoveInstruction = MoveInstruction {
            steps: 0,
            direction: Direction::None,
        };
        parse(&mut move_instruction, instruction);

        //print!("Process {:?} ### ", move_instruction);
        process(&mut position, &move_instruction);
        //println!("Now in {:?}", position);

        // work out path covered
        if !has_found_actual_location {
            let delta_x = i32::abs(position.coordinate.x - old_position.x);
            let direction = i32::signum(position.coordinate.x - old_position.x);
            for i in 1..delta_x + 1 {
                let visited_path = Coordinate {
                    x: old_position.x + direction * i,
                    y: old_position.y,
                };

                if !has_found_actual_location && visited_places.contains(&visited_path) {
                    actual_location = visited_path.clone();
                    has_found_actual_location = true;
                }
                visited_places.insert(visited_path.clone());
            }
        }

        if !has_found_actual_location {
            let delta_y = i32::abs(position.coordinate.y - old_position.y);
            let direction = i32::signum(position.coordinate.y - old_position.y);
            for i in 1..delta_y + 1 {
                let visited_path = Coordinate {
                    x: old_position.x,
                    y: old_position.y + direction * i,
                };

                if !has_found_actual_location && visited_places.contains(&visited_path) {
                    actual_location = visited_path.clone();
                    has_found_actual_location = true;
                }
                visited_places.insert(visited_path.clone());
            }
        }

        old_position = position.coordinate.clone();
    }
    let shortest_path = get_shortest_path(&position.coordinate);
    println!("{}", shortest_path);
    let mut output1 = File::create(&"output1").expect("Unable to create");
    output1
        .write_all(shortest_path.to_string().as_bytes())
        .expect("Unable to write");

    let actual_shortest_path = get_shortest_path(&actual_location);
    println!("{:?}", actual_shortest_path);
    let mut output2 = File::create(&"output2").expect("Unable to create");
    output2
        .write_all(actual_shortest_path.to_string().as_bytes())
        .expect("Unable to write");
}
