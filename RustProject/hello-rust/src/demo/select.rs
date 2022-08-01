use futures::{
    future::FutureExt, // for `.fuse()`
    pin_mut,
    select,
};

use std::{thread, time};

async fn task1() {
	thread::sleep(time::Duration::from_seconds(5));
}

async fn task2() {
	thread::sleep(time::Duration::from_seconds(3));
}

async fn race_task() {
	let t1 = task1().fuse();
	let t2 = task2().fuse();

	select! {
		() = t1 => println!("t1 is done!"),
		() = t2 => println!("t2 is done!"),
	}
}

fn main() {
	race_task();
}
