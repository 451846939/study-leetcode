/*
fn main() {
    // let boss = tokio::runtime::Builder::new_multi_thread().max_threads(1).build().unwrap();
    let works = tokio::runtime::Runtime::new().unwrap();
    works.block_on(async{
        let listener = TcpListener::bind("127.0.0.1:8083").await.unwrap();
        loop {
            let x = listener.poll_accept().await.unwrap();
            works.spawn(async move{
                println!("{:?}", x.0)
            }).await.unwrap()
        }
    })
}
*/

impl Solution {
    pub fn remove_stones(stones: Vec<Vec<i32>>) -> i32 {
        let n = stones.len();
        let mut edge:Vec<Vec<usize>>=vec![vec![]; n];
        for i in 0..n {
            for j in 0..n {
                if stones[i][0]==stones[j][0]||stones[i][1]==stones[j][1] {
                    edge[i].push(j)
                }
            }
        }
        fn dfs(edge: &Vec<Vec<usize>>,i:usize,vis:&mut Vec<bool>){
            vis[i]=true;
            for y in &edge[i] {
                if !vis[*y] {
                    dfs(edge,*y,vis)
                }
            }
        }

        let mut num =0;
        let mut vis=vec![false;n];
        for i in 0..n {
            if !vis[i] {
                num+=1;
                dfs(&edge,i,&mut vis)
            }
        }
        (n - num) as i32
    }
}