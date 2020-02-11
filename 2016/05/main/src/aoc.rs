pub mod aoc {
    use md5::{Digest, Md5};

    fn find_next_hash(key: &str, current_hash: &mut u64) -> String {
        let mut hasher = Md5::new();

        loop {
            *current_hash += 1;

            //let next_key = key.to_owned() + &current_hash.to_string();
            hasher.input(format!("{}{}", key, current_hash));
            let hash_candidate = hasher.result_reset();
            let hash_candidate_str = format!("{:x}", hash_candidate);
            if hash_candidate_str.starts_with("00000") {
                return hash_candidate_str;
            }
        }
    }

    pub fn find_unordered_password(input: &str, n_iterations: u64) -> String {
        let mut current_index: u64 = 0;
        let mut current_hash: String;

        let mut solution = String::new();
        for _i in 0..n_iterations {
            current_hash = find_next_hash(input, &mut current_index);
            solution.push(current_hash.chars().nth(5).unwrap());
        }
        solution
    }

    pub fn find_ordered_password(input: &str, n_iterations: u64) -> String {
        let mut current_index: u64 = 0;
        let mut current_hash: String;

        let mut solution = String::with_capacity(n_iterations as usize);
        for _i in 0..n_iterations {
            solution.push_str("?");
        }

        loop {
            current_hash = find_next_hash(input, &mut current_index);

            let position_index: usize;
            match current_hash[5..6].parse() {
                Ok(x) => position_index = x,
                _ => {
                    //println!("hash({})[{}] cannot parse index", current_hash, &current_hash[5..6]);
                    continue;
                }
            }
            if position_index >= solution.len() {
                //println!("{} | {}[{}] <- cannot use", current_index, current_hash, position_index);
                continue;
            }

            unsafe {
                let bytes = solution.as_bytes_mut();
                if bytes[position_index] != b'?' {
                    continue;
                }
                bytes[position_index] = current_hash[6..7].chars().next().unwrap() as u8;
            }

            //println!("{} | {}", current_index, current_hash);
            if solution.find("?").unwrap_or(n_iterations as usize) == n_iterations as usize {
                return solution;
            }
        }
    }
}
