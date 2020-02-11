extern crate main;

#[cfg(test)]
mod tests {
    #[test]
    fn case_study_1() {
        const N_ITERATIONS: u64 = 8;
        let solution = main::aoc::aoc::find_unordered_password("abc", N_ITERATIONS);
        assert_eq!(solution, "18f47a30")
    }

    #[test]
    fn case_study_2() {
        const N_ITERATIONS: u64 = 8;
        let solution = main::aoc::aoc::find_ordered_password("abc", N_ITERATIONS);
        assert_eq!(solution, "05ace8e3")
    }
}
