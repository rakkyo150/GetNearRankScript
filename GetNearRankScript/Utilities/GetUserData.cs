using Newtonsoft.Json;

internal class GetUserData
{ 
    public async Task<int> GetYourCountryRank()
    {
        int yourCountryRank;

        string yourBasicPlayerDataEndpoint = $"https://scoresaber.com/api/player/{Config.Instance.YourId}/basic";

        Console.WriteLine("Start GetYourCountryRankPageNumber");

        HttpClient client = new HttpClient();
        var response = await client.GetAsync(yourBasicPlayerDataEndpoint);
        string jsonString = await response.Content.ReadAsStringAsync();

        dynamic jsonDynamic = JsonConvert.DeserializeObject(jsonString);

        string yourCountryRankStr = JsonConvert.SerializeObject(jsonDynamic["countryRank"]);
        yourCountryRank = int.Parse(yourCountryRankStr);

        Console.WriteLine("Your Local Rank " + yourCountryRank);

        return yourCountryRank;

    }

    public async Task<HashSet<string>> GetLocalTargetedId(int yourCountryRank)
    {
        int yourCountryRankPageNumber = 1 + (yourCountryRank - 1) / 50;

        string basePageEndpoint = $"https://scoresaber.com/api/players?page={yourCountryRankPageNumber}&countries=jp";
        string lowerRankPageEndpoint = $"https://scoresaber.com/api/players?page={yourCountryRankPageNumber + 1}&countries=jp";
        string higherRankPageEndpoint = $"https://scoresaber.com/api/players?page={yourCountryRankPageNumber - 1}&countries=jp";

        int lowRank;
        int highRank;
        bool otherPage = false;
        int branchRank = 0;
        HashSet<string> idHashSet = new HashSet<string>();

        Console.WriteLine("Start GetLocalRankData");

        Dictionary<string, string> result = await GetCountryRankData(basePageEndpoint);

        lowRank = yourCountryRank + Config.Instance.RankRange;
        highRank = yourCountryRank - Config.Instance.RankRange;

        // トッププレイヤー用
        if (highRank <= 0)
        {
            highRank = 1;
        }

        // ページを跨ぐ場合
        for (int i = 0; highRank + i < lowRank; i++)
        {
            if ((highRank + i) % 50 == 0)
            {
                otherPage = true;
                branchRank = highRank + i;
            }
        }

        if (otherPage)
        {
            if (yourCountryRank > branchRank)
            {
                Dictionary<string, string> resultSecond = await GetCountryRankData(higherRankPageEndpoint);
                foreach (var resultPair in resultSecond)
                {
                    result.Add(resultPair.Key, resultPair.Value);
                }
            }
            else
            {
                Dictionary<string, string> resultSecond = await GetCountryRankData(lowerRankPageEndpoint);
                foreach (var resultPair in resultSecond)
                {
                    result.Add(resultPair.Key, resultPair.Value);
                }
            }
        }

        for (int i = 0; lowRank - i > yourCountryRank; i++)
        {
            idHashSet.Add(result[(lowRank - i).ToString()]);
            idHashSet.Add(result[(highRank + i).ToString()]);
        }


        // トッププレイヤー用
        if (idHashSet.Contains(result[yourCountryRank.ToString()]))
        {
            idHashSet.Remove(result[yourCountryRank.ToString()]);
        }

        return idHashSet;
    }

    public async Task<Dictionary<Tuple<string, string>, string>> GetPlayResult(string id, int pageRange)
    {
        // Key=>(songHash,difficulty),Value=>pp

        int topScoresPageNumber = 1;
        var playScores = new Dictionary<Tuple<string, string>, string>();

        for (int i = 0; i + topScoresPageNumber <= pageRange; i++)
        {
            string playerScoresEndpoint = $"https://scoresaber.com/api/player/{id}/scores?page={i + topScoresPageNumber}";

            HttpClient client = new HttpClient();
            var response = await client.GetAsync(playerScoresEndpoint);
            string jsonString = await response.Content.ReadAsStringAsync();

            dynamic jsonDynamic = JsonConvert.DeserializeObject(jsonString);

            foreach (var jsonScores in jsonDynamic["playerScores"])
            {
                string songHash = JsonConvert.SerializeObject(jsonScores["leaderboard"]["songHash"]).Replace("\"", "");
                string difficulty = JsonConvert.SerializeObject(jsonScores["leaderboard"]["difficulty"]["difficultyRaw"]);
                string pp = JsonConvert.SerializeObject(jsonScores["score"]["pp"]);
                var hashAndDifficulty = new Tuple<string, string>(songHash, difficulty);
                playScores.Add(hashAndDifficulty, pp);
            }
        }

        if (playScores == null)
        {
            Console.WriteLine($"No {id}'s Play Scores");
        }

        return playScores;
    }

    public async Task<Dictionary<string, string>> GetCountryRankData(string endpoint)
    {
        // Key=>rank Value=>id
        var countryRankAndId = new Dictionary<string, string>();

        HttpClient client = new HttpClient();
        var response = await client.GetAsync(endpoint);
        string jsonStr = await response.Content.ReadAsStringAsync();

        dynamic jsonDynamic = JsonConvert.DeserializeObject(jsonStr);

        foreach (var jd in jsonDynamic["players"])
        {
            string rank = JsonConvert.SerializeObject(jd["countryRank"]);
            string id = JsonConvert.SerializeObject(jd["id"]).Replace($"\"", "");

            countryRankAndId.Add(rank, id);
        }

        if (countryRankAndId == null)
        {
            Console.WriteLine($"No Country Rank and Id at {endpoint}");
        }

        return countryRankAndId;
    }
}
