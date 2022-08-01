use std::collections::HashMap;

fn __main__() {
    let mut mp = HashMap::new();
    let key = String::from("name");
    let val = String::from("zhang3");

    mp.insert(&key, val);
    mp.entry(&key).or_insert(String::from("li4"));
    match mp.get(&key) {
        Some(val) => println!("{}", val),
        _ => println!("done have this value"),
    }

    //println!("{}, {}", key, val);

    let mut scores = HashMap::new();
    scores.insert(String::from("Blue"), 10);
    scores.insert(String::from("Yellow"), 10);

    let team_name = String::from("Blue");
    let mv = scores.get(&team_name);
    match mv {
        Some(&val) => println!("{}", val),
        _ => println!("done have this value"),
    }
}
