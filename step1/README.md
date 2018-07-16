# 概要
ただただ、データを格納するだけのコード。  
Nameをプライマリーキーに設定しているため、複数回格納を実行してもデータが増えることはない。  
ただ、created_timeは新しい値で更新される。  

# 初期テーブル定義

    tablename = Test
    Name S [partition key]
