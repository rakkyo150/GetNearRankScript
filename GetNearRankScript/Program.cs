using System.Reflection;

// Environment.CurrentDirectoryはMacだとUsername直下になる
string configFile = Path.Combine(AppDomain.CurrentDomain.BaseDirectory,"Config.json");

ConfigManager _configManager=new ConfigManager();
Progress<string> progress = new Progress<string>(onProgressChanged);
GetUserData _getUsersData = new GetUserData();
PlaylistMaker _playlistMaker = new PlaylistMaker();

Console.WriteLine("Start");

if (!File.Exists(configFile))
{
    Console.WriteLine("Make a Config File");
    _configManager.MakeConfigFile(configFile);
}

_configManager.LoadConfigFile(configFile);


await GeneratePlaylist(progress);



async Task GeneratePlaylist(IProgress<string> iProgress)
{
    try
    {
        var othersPlayResults = new List<Dictionary<Tuple<string, string>, string>>();


        iProgress.Report("Getting Your Local Rank");
        int yourCountryRank = await _getUsersData.GetYourCountryRank();


        iProgress.Report("Getting Rivals' ID");
        var targetedIdList = await _getUsersData.GetLocalTargetedId(yourCountryRank);


        iProgress.Report("Getting Your Play Results");
        var yourPlayResult = await _getUsersData.GetPlayResult(Config.Instance.YourId, Config.Instance.YourPageRange);

        iProgress.Report($"Getting Rivals' Play Results");
        foreach (string targetedId in targetedIdList)
        {
            Console.WriteLine("Targeted Id " + targetedId);
            var otherPlayResult = await _getUsersData.GetPlayResult(targetedId, Config.Instance.OthersPageRange);
            othersPlayResults.Add(otherPlayResult);
        }

        iProgress.Report("Making Lower PP Map List");
        var hashAndDifficultyList = _playlistMaker.MakeLowerPPMapList(othersPlayResults, yourPlayResult);

        iProgress.Report("Making Playllist");
        _playlistMaker.MakePlaylist(hashAndDifficultyList);

        iProgress.Report("Success!");

        iProgress.Report("Press enter to close");
        Console.ReadLine();
    }
    catch (Exception e)
    {
        // iProgress.ReportだとConsole.WriteLineが先に来ちゃう
        Console.WriteLine("Error: " + e.Message);
        Console.WriteLine("From the beginning again!");
        _configManager.MakeConfigFile(configFile);
        _configManager.LoadConfigFile(configFile);
        await GeneratePlaylist(progress);
    }
}

void onProgressChanged(string debug)
{
    Console.WriteLine(debug);
}