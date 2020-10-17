package bertyprotocol

import "html/template"

var templateAuthTokenServerRedirect = template.Must(template.New("redirect").Parse(`<!DOCTYPE html>
<html lang="en-GB">
  <head>
    <meta name="viewport" content="width=device-width, initial-scale=1.0" />
    <title>Redirect</title>
    <script>
      location.href = "{{.URL}}";
    </script>
  </head>
  <body>
    <div
      style="display: flex; align-items: center; justify-content: center; height: 100vh"
    >
      <a
        style="text-decoration: none; font-family: Verdana, Geneva, sans-serif; font-weight: 600; color: white; border-radius: 10px; background-color: #404148; padding: 10px 17px; display: flex; align-items: center; justify-content: center;"
        href="{{.URL}}"
      >
        REDIRECT
      </a>
    </div>
  </body>
</html>
`))

// nolint:gosec
var templateAuthTokenServerAuthorizeButton = `<!DOCTYPE html>
<html lang="en-GB">
  <head>
    <meta name="viewport" content="width=device-width, initial-scale=1.0" />
    <title>Connect</title>
  </head>
  <body>
    <form
      method="POST"
      style="display: flex;  align-items: center; justify-content: center; height: 100vh"
    >
      <button
        style="border: none; border-radius: 10px; background-color: #525BEC; padding: 10px 17px; display: flex; align-items: center; justify-content: center;"
        type="submit"
      >
        <img
          alt="Icon"
          style="margin-right: 17px"
          src="data:image/svg+xml;base64,PD94bWwgdmVyc2lvbj0iMS4wIiBlbmNvZGluZz0iVVRGLTgiPz4KPHN2ZyB3aWR0aD0iMjlweCIgaGVpZ2h0PSIyN3B4IiB2aWV3Qm94PSIwIDAgMjkgMjciIHZlcnNpb249IjEuMSIgeG1sbnM9Imh0dHA6Ly93d3cudzMub3JnLzIwMDAvc3ZnIiB4bWxuczp4bGluaz0iaHR0cDovL3d3dy53My5vcmcvMTk5OS94bGluayI+CiAgICA8dGl0bGU+c2VydmVyXzFfZm9yZ2luZ19ub2RlPC90aXRsZT4KICAgIDxkZWZzPgogICAgICAgIDxmaWx0ZXIgaWQ9ImZpbHRlci0xIj4KICAgICAgICAgICAgPGZlQ29sb3JNYXRyaXggaW49IlNvdXJjZUdyYXBoaWMiIHR5cGU9Im1hdHJpeCIgdmFsdWVzPSIwIDAgMCAwIDEuMDAwMDAwIDAgMCAwIDAgMS4wMDAwMDAgMCAwIDAgMCAxLjAwMDAwMCAwIDAgMCAxLjAwMDAwMCAwIj48L2ZlQ29sb3JNYXRyaXg+CiAgICAgICAgPC9maWx0ZXI+CiAgICA8L2RlZnM+CiAgICA8ZyBpZD0iYmVydHkiIHN0cm9rZT0ibm9uZSIgc3Ryb2tlLXdpZHRoPSIxIiBmaWxsPSJub25lIiBmaWxsLXJ1bGU9ImV2ZW5vZGQiPgogICAgICAgIDxnIGlkPSJleHRlcm5hbHNlcnZpY2UvY3JlYXRldG9rZW4tY29weS0yIiB0cmFuc2Zvcm09InRyYW5zbGF0ZSgtOTQuMDAwMDAwLCAtMjU2LjAwMDAwMCkiPgogICAgICAgICAgICA8ZyBpZD0ic2VydmVyLTEtZm9yZ2luZy1ub2RlIiB0cmFuc2Zvcm09InRyYW5zbGF0ZSg3Ny4wMDAwMDAsIDI0NS4wMDAwMDApIiBmaWx0ZXI9InVybCgjZmlsdGVyLTEpIj4KICAgICAgICAgICAgICAgIDxnIHRyYW5zZm9ybT0idHJhbnNsYXRlKDE3LjAwMDAwMCwgMTEuMDAwMDAwKSI+CiAgICAgICAgICAgICAgICAgICAgPGltYWdlIGlkPSJzZXJ2ZXJfMV9mb3JnaW5nX25vZGUiIHg9IjAiIHk9IjAiIHdpZHRoPSIyOSIgaGVpZ2h0PSIyNyIgeGxpbms6aHJlZj0iZGF0YTppbWFnZS9wbmc7YmFzZTY0LGlWQk9SdzBLR2dvQUFBQU5TVWhFVWdBQUFKWUFBQUNYQ0FRQUFBQmRWR3JNQUFBQUJHZEJUVUVBQUxHTjVmSUFLUUFBQU5KcFEwTlFTVU5ESUZCeWIyWnBiR1VBQUJpVlkyQmdyRWdzS01oaEVtQmd5TTByS1hJUGNveU1pSXhTWUwvS3dNN0F5QUFHaWNuRkJZNEJBVDRNT01HM2F4QzFsM1ZCWnVGV2h4V3dwS1FXSndQcExVQmNtbHhRVk1MQXdLZ0RaS3VYbHhTQTJDRkF0a2gyU0pBemtKMEJaUE5CMVlPQXRITmlUbVpTVVdKSmFvcUNlMUZpcFlKemZrNStVWEZCWW5JcWlhNGdBcFNrVnBTQWFPZjhnc3FpelBTTUVnVkhvRzlUZ1hibUZwU1dwQmJwS0hqbUplc3hNSURDRDZMamN5QTRYQmpGemlTWEZwVkJqV0ZrTW1aZ0FBQVJYVFRGSFhuVGtRQUFBR3hsV0VsbVRVMEFLZ0FBQUFnQUJBRWFBQVVBQUFBQkFBQUFQZ0ViQUFVQUFBQUJBQUFBUmdFb0FBTUFBQUFCQUFJQUFJZHBBQVFBQUFBQkFBQUFUZ0FBQUFBQUFBRXNBQUFBQVFBQUFTd0FBQUFCQUFLZ0FnQUVBQUFBQVFBQUFKYWdBd0FFQUFBQUFRQUFBSmNBQUFBQW1EMVlXZ0FBQUFsd1NGbHpBQUF1SXdBQUxpTUJlS1UvZGdBQUZJWkpSRUZVZUFIdG5RdWNYRlY5eC8vejJGZVdaRjhraGdSQ2lER0VKSVFvRW9nQlNTSzBDa0xsVXhWRmZOTnFiYXVmcWxYN1VteHJhMzE4eE1mSDB2WWpXaFJLb1IvOFNGUWc5S0dnUWlJa2tNUVFROTRodThubXNkbmQ3Q016c3pQOW5udm56RDMzenIyek01UFpuWm1kK2U4SDVqN09QZmVjMy8yZi8rdjh6MGxJRkRYSmxmSk9XU096ckxOYSsxK0lEZy9KZm5sWS9rc081ZXA4bEpzUm1TK3ZseFd5VU5SWmJWS0hOTU13KzJSQStuTUJFSkZYeWQyU3F2K0J3QTU1QjZ3VFNBcXFOOG9ucFNHd1JDM2Q2SlRaOG5PNEsrWGY2YkRjS091a3hmOW16VjJOeUNKNWN6QWFZVmtsUzJzT2xPQU9ud01lelVHM3c5SWgwNE51MXVEMUJtbVhjRkMvdzVMa3IwNGFnUlJvQkVnc0FVV2xCK3ZrSUpBRGpVQ1djNTZ1SDJrRTZtQnBKUEw0cllPVkIwaTZTQjBzalVRZXYzV3c4Z0JKRjZtRHBaSEk0N2NPVmg0ZzZTSjFzRFFTZWZ6V3djb0RKRjJrRHBaR0lvL2Y0TmpvUzNKRVlzRk9wVS9kWTNLT0xNWVJuYkxrRDlaUitWOTVRbDZXTXdXQmxaQTJlWTNjTEplTGltcFBRZklES3lZL2xMK1FVMFgxOWlmeWpIeExYcGtyT0Z0VXZSWHhrSi9NMml6cmM0ZnRjN2I4ZWJsZlR1UXNVYlUzL2NEYUxqdlBJbXpUTDc4Z2lqMGx5UStzdUNUT3FxL3hxUnBPOUFQck5laTA0cWxGMWt4VmplZ240SytVR3hEVG96SldNR0locHRSV01MYzlzK0FucStJQlA3QkUzazEzNzVjRE1vTHBrTDhaa01UT3VrN2VoUzdNLzVtcUFFazMwaCtzVHZrOVdTNm40YTFDdXAyQ3I4NlRDM1RWVSs4M2lrWGtKN2RhWmRuVTYyd2VQUXI1b3BGK01JeVZIcytqa2xvcGtnU1B3UG1kTUtrMlBiV0NSQjc5SEphRHdZb3RqQSs0Slk5S2FxWElvR3hBclFWUVJJNlN5TFdpUnRQWXZLQ2N3UWI0bmd4N0wrdnpDUGJVVVRrbW5lUThCQ1pFNk1KVC9IZS8vSWY4RzhNd2tMUnBzRnJlUnY1SVI5cFJTWklEMkE2QU9yRXJBYUQ5bWJQQXlpcjZSZ285Tnh0TFVPditZVG1PY2FSNnFEQTRRMEJxZzl3SDIrUWdEWlo2NUZ4TVVidXFZYXI5ZmZsQXhtM3BsYjhsNDdJclJ6MlZmeXNHVUg4anZ5UFQwazE5V3Y1Wm51S2E2bmtLT2RYRGY0RjYwSDdHTVVwVG9Ib3NiWVNtaUdZZE5OenBVYXo1SGlLbjFVd3BrdFNPRzVxdVQxN0NFdERNTWc1TVhyRHNjLzNRa01mNlVocEMzNnRXeUVhTXp5OEluSUw3cERuckhISkw1eU9sYkRyRGtMekdTQmRzazdjZ3o5cXFGU1dyM1hFRzRLWFNtT25EQXJrTkwwVXBOY1ZkbzR5YlBXUXI1eVJWTUl6aDhDWUUvTlUxbkFNNGhvQmZMeitRRjRBdEo1MkgrRDRGaTZxY3QxcjlTeUxOWW1RcVg1OExxWWljVDFqbE0yaTZRb0l4dVdxc3puc2hCbU9FbUVtYlBJbFJHaUNkSTNLci9DNXJDK3FrRUlnZ3RROHlBeEVRV2dqTFdpSlhkZElJdE9aYUZSQkc3MVczbHRQZExNMXZBN3lsYmZ5c0dsVzJjcDN5UktDZTJ1MEZLZ2Z6QkxLY3Q0NzZ1VEpJNjVRM0FyVURWcjRyU1hJTVErMGI1bzF1MVJaVUtRbitNMW51THRrT3R2dGErbXhxZzdWZlhpUzQxQ3NuWmRDYXRZa1FIR2pIRTU1TDBPQVNYNU9wZ2VzM3lvUCtOc0xVQkd1TUJieDdaWnZzSXBMd01nbFFRd2FqS0N2OUZUaDVDMWxudVpSWXkyempucExoRjhuSENRaytMSDJ1NjliSjFBTnJpRERsUnZreFVkREQyZDNseXBnVjVOek9rVnFZdWc3UGVCRzg1c2p1aUt3RXpFYjVUL2d4aXg0RCs2a1NhMGpDUXcvU2ZUMTNrTlZabnd1ejVjL2dRVy9NWmI5ODFBQXcvVmhFYmdmWEdUNlZWT09sVGZJNXVZY2hHT0FJKzNicE5JN3pMNGxpbVlGQkFaRUxtYzdZNGMwMG15cWNkVnErUS9peTJPbThCWElIazgzdUViYWRrS2hMVExsT1hNZ2ZSc1Fsc2xuUlZjWjlrcVNwYzhzVWJlMUd4dHlOUUMrVzlxSVFUc3FmeUxWR2o1ZktKMUFPbTV3cS9jRTZEY3FQOHZob1FhTS9RWngrTlY5amp2RkM1MDBUZDVSaTF1WWV1Y3MzN2JlSnlIc1Q0cnFaLzZmb1R4d1RZcFNwaWpNK3pYbVlHdTZrQjg0ZUYyOGcyMnkzS2VqOWh1SDlWbkoybUU0WCtqZU52TDhlRHp1N21idjBaNzN5ZDh4NHF0a0VMN1hKRytRdjVmdUVpMy9EeE5jdUJQbFR4Tm4vbkJpZWYxZ3FERlRQRWw1MjJ0Z0x0eG1VRGRaV2VidHh2OUREdWZKMXZwRHp1b2srR3BKL1JkVm4wOFdBc2dHQWpwQTdmU1k5djZEaTdBTjh6RjN5bUh4S0ZtUS9CQTllejEyenpldk5ETnRzc0w2THdWWThOYkJkeTI3WDY4eFhsLzc0RVo5TlBNNlhqekZYY3lobksxN0c4THlEcUx1WHdsdzFuendwWDNDS1pJUDFMWm5uM0M3NHFCR2Q5TnVjelN3bFlEdmxKazhMUXl5SitTWnl4bnlMNHFZRFJOZDdySDFtbkRzOVNMb3IzUnFQMmhwNFh1MUhvLysyOHprczJlNW5adzNLVnFvdWxxYVQ2N3h1a25SaVRMNGtEeGhUOHFxaks1Rmd0NlVuaUJQNGhMMnltV1VNLzRmYzJvVE1Pb2paMnNDZ2pGcEs2Qnk1Q3YxOWdLRnFabVluZ2ZxVnh1aHFZTDVubSsweVpYTlduQVNLNG1rSnpvWXBJdlgzS2YxdkhHTnlwVXV3UitRS3JqbHYzeW1mQlk1ejBZaEtHemJ6TjQyenRjQ3AxcERZZnpIUythNzFkRGZLTUhabVVlUHl2Q3hSSmZ4TWg2aThYMXJsUi9oV28rRHZwMlU4ZFZ1bmFtT1NWb2JnYmN3V09jclhyMlNwcmlYa1hyU2NHWDlhTHY4RWZIYWYrdVM3U0tYZGVJTHU3V09HNGJPZFNMUmI1YjFXWGxBRGJmNFkvUGVpMGF3RWtEK083TFVwU2lUaVVyUnAzQThzd2ZQK0kza3RPaTBPVlBtRGxlTExYV0ovQStQRkUzV1l3Qm5aNElvbm5JY1dYNXQrM1ZiNUYza0VyOWVQNHNndTlYY1l1TlEwWURPd0hHY2RuTGs4YTRjOGhGN1VQbVlqdkxkSlpVSnM0Q0hOa3RYMHEzTEdkSzZWRGNsU09tajM0SG1XUFl4UEVYa2YvSkt3bmpsdGpTYnptV1Z3clgxUDFibE5UZXlITFgvYkxGUXR4NGV4bGR3THNyb3hPZnZodGQzeWJRelI4V2tNNWZCRlJMNGF5RTFFR1JhN0hqbUpMQnZNWEhrVml5RmFJdVRDdGZzYWRabHlGWHF3aFVXZzdveVhHQU5sRjRQekhnWmdmbkdIQkViT0lnUkhBMHpUaFJqZmJQUlZlYnJYWmJSNkJON2FHeUVxMkNxWEdZV3E0M0NVNy82SXA2a3B1RXJGUjNkNFFQUVVjNTNHcFJ1NWRTSFhJbkJZTjBaRGEvcnBNUUMvbWZDVlRTRmlyajBxdFhzVWhkdmhxcUx5VDdyUjFyLzJhV2JNV25Ia2N5UHdVZzlRWFdWcDBGbllWcGZqVGM3QTdsS1pOQ09rekRpTHRvN0tpUWhmNHdpKzB6eUVwWk1WRjFoenhkellDVmd2ZVZxVHI5NzJQQVl2cmJBRVVRdlFYTUhmbFh5R1BSUktFblpla2tIbGxKeFNwc05KUnY5ZUlxYXZRM3JaTm9sS2cyN2dUNy9lRG00NGNXcnY2eWJqUElRUWRneWRVM3hpazhKMDl4WHd3aDVzY2RQeU1zc0VIZTlCV2lrVFF0UDVXUFdLUWd6TEFTdWZXWjIxeVN6NzlYRWt3RWFLektLQWV0VUlJWTgzeVZzejRxMVB2b0h1MFJtbjZ0SEpKYVc4cDJFSlhaRjU3V25QN2dEejVLOVFWUWZscS9MZkRNVENTUG1OSmgxTTIyY3AyTWhSSWEzU3JyL1ZNS1BVL0ZadGhQQnV5b0ExQkdOdU5Pc3J3M0ZZL3NCNDY2akxIRzFoVzlxMzhlM1B4MTdmVWpCWUE2NTVvRUg4emVmU2IxS3pqWm9JSTJxdzlDWDlHM01wMzZUclRKZVozTjhrb3NHaGhOR2laaXp4MjlMaHZJdU5IR3VuZE80akpXUTBEZUZDM1pmaFduTnBmRmlpR3F4bTdJeFptUlVWbzR6L0ZSblJKb2pBbFhoUDVSeUdBcGViSVQ0VndiV3BTZGJJcDlNdHo3bGFVT1BoODZ1bFhJem94SmN5VUNtVFhVdHR0UW9nYVlQVmpPL3pMbmJ1UGpjakhFTUlVMGM3ZGlJUlBtMDg2UE8rQ2I4VXdraDBxQ0Z6ZGltQjMvbnBHMkhFc1ZuS0taLzc2QVNUTSszMDd3WFBRcWRtUTZVa0pLYkFhZ09vanlLanBnZldHQzZDdVFNcks4bU5sclNXbXNFYW83V1pHbFZHZnpIOC81emNpUk0reW1UYS8yVHFVdHF3RFEyc2FWUk9SOUY3cTVtVFhhaXZWY252REVhQm9qYmlIS1k3UFIxYjZSbTBlV0YwQkY5eU01N21iMXhxSTJXdEl0TTFEY2p4TUJiclc2b09Lb0dEN05EM0NGSGQ0MWlIMmhjTUkxQXUwUDByNExlUEtOY3podU9zSGcxajJ6dWo3WVFjaW1MU3Y3cUFTaXVsNk96MEJ6NU81RUJGMFVhd3RTKzJKTXhheTdiWEl2dHMydHNCV0k3YzdwVjl5dkt0TnI5UUFkREpjTE5wSjBId2R4Qy91aXR0V3M1QnJKUWlkeVBLVkptWmpuUlVnUlUxSkg3Ni9WWHdFNklyRjFpUlRHVXVEUEhmVCtWcHE5MGhSSFVoODU2dnhsajRIczZldzBOMjk2ZmpMNmlsbTVvT0tIL3g4U3FObE82VDl4aTZTblhwZG14Qk8xTDZ0S0VoZFdmOWZ4ZGpndllqOHpZaTY5eEd4M3pnVjJ0YjdiK2phTjF3OWViQnoyUUxHUGR3ZTVMY0xEc1FzQks3OEdvUGxObGdoWmhuK0JSY09BTWpZU1U1Z09ZMFM1aWdvUGx2RG13azFFRGQyVk5oR3MxSy8rMW1BYW1lVXJDaHVNd3lBRlM3RmEvY0RBenUrdzVnSWJ5U3F6SXgrekZpckd0Y1lNMlJMMmU0U3RYMlNkdCtxMTZ3aGhEcXRnR2hRWmhHOHV5K2RDZGpUTUova1hrcWYrcVVQd1hZMFhUWkhzQndqQVQxeEExSUtJZForaGpXbG50VnZXQWxpZkxlNk1HaUJiMzRZcWFidmV5SytZK0VtaHozT29wYWVDT1pOVDloTWJrR28xZitHcnZOcEM3NW1oVVhzMHVNc0dHa0JicnRHNW9GcStjNFJCZmZ6WVNwaW1wcUdpR3BMY1dzNSt1c0N6UGhpSlZFVlBjeVM5aVB4b3d5TUdmaVNTNUsyLytxMEU0bVk3K1BhakRwRnVTaGR0UUZ2L0ZlSGI2cVhzNVMzMzJFMmNQc1hLdnJyUndhWndKZTg1RDM5NWo4aXQwcjNGcFFNRzkva2VFNkZiNzVrVE5FcXhzczlVOVczZTVqS1o2TEZOcU0vUlZET0hzaFVnSTdEc3o3RWVJWG13ekZjUWpPZThDYXJ0QlBiV1dpUDZSTFZUdFljZlRlYXFjNzZXNkZHSEJMNUNNa2Uvb2wxdlhEVVo5bFlxSWpDK1l1c3JGVWRvU0dLa1dxWE1aUnIyYVpaZU1TSlIvckMvSjV3blltcFpocTJNR0V3eStaV1ZoQXVIa1dRNmtSRUFaeHV3OGo1VlRjL1lqNWdIWGNCYndmTU9TWjhQeTk4Rm1hcWg4c2xRbDBMWk1aMDdFWVk3cGI2Vittcndqb0NSelVoZXVpOHJLR0VOZmV2QnI5MEh6QzAzZTQvTUhEQkFOdEo4b3FFd3pXU1VaMU1vdTlkY1YrdnltYTNlWXdyVitSQ2J0MlBZQk1JM1RudjB0Ukh4RGxwaWE0N3lOTWlaaldWaitKQUkrYWFXNytZQ1V4N2RiRHJJVnR2Wm5BeXIyVytGaDUvazI3MThwWFNBajVUaURmQklPbFppUlg0ZmlzZFRsSWd3QjF0OXVrQ01HOHk5S1RpazUxRzVBQ083RkxDb3NLS2M1cWwxdmtUdGYzY1dxZDZLTVVSdW9XWk14REpqZms4ZEtGY05SYlNjVjFmK1QxckE2ei9FR3pobXh0dUJzaFZ6ek5rMy9IQUhTMHlXUWZiWU1mM3U2U1BNRjlhU0FNOHdtcytXTlo3WDJJckljczhodUdHOUVCeFZNUHlZblhvTGpMUmNzWUtWZWpHN2NnUmw1bUdEazVWazZMR3BCd3M3SGtMd09zMTJjWnRRUEE5N1hNUkt2elZKYWRvVzZkY0pTbFVUTGZ3NUNWWHBsdjZZa3BwNVpkSm1VN211d0ZUTStqbUJIMmRLbmFzclVGemFpV1psNk9uREpuSXUyV0tKUGpJZmw3ejRSK3VwVituRFdQU2c0VjNZdEdBcjVlQjZMb3lzN2l3VEJKczhzWVhxTmtMSnhFRzZva29nWU1pQzdzcURhMHZIKzI3Qmp3Zmc0cnpKZjh3RnFEQ243V3QzUStGOXRadmVQMjRmTjVhaUxLMkhDMEVybWF5elRYR0s5UU94bEZjNXBEWXd6YndGMk8vTUJxUThBblNiTHZ4Y2h6Zk8veHU1T0V4WmVqVjRwZjlUZitPNG9yRWM2S3NBZlZvOVNSR1RGMWxmTURTOGh6K0R4WnpOMFlENFdCcFhJaVZybnFyN1lUeFh1QnBKZ3k0MU1icGJvWVRIWHlJR0RtaVhodTFlU3BIK05rZ0Fpaktmd3NrVXlCR2p1STRYcmJNMFErSFkvS3p4aWwzaENZVDhFYXVUUU1IazVxbTZmVEVVejlNWVJ5cStkNmJaN0dzZnkvd1ZnTDhJa2orSEc5S012Ri9KZEREOVFBZGlvTllCTkxscDhMZ2txdEt4Q01zRzFFRFRzbmZkVjhaWDJCWTB4NGZSbXZXQm12QWFTbGZ3Y3pZeGNSckxkVHV3TUtUOERsTXdSMWJqREV3Q0htVXA2ZVpEZGNZVEJNc0hrL1Vid2NVRTFBN3d1dWNnYnp4c2RoZmZ2dldmTEZLcFlLc2RBbnBoT0tseDJCcXVaVktwYktENWJidFZWbkZVdmxCNnRpb2NsdVdDV0FwWlZNZHVzcTdFcjV3VkpUNlk2Y0dxdGtmVlIrc0p5RkpZcVBJcFZzR3BjZnJEbGtURFZseGxzSDZXbmxiMU9tT2U2RGNzcUxNRGJWTGVSMXppY21yZ0dLV1hOL1Q3SCt1ZmlOdzl3OW5CSm4xNURrK0dzcjNxME5VdWYzSUZPL2Yxd2hrZnl5Z3oyRFNQM2psakhxQU9ROTZtWkM2cEt5dDdUc0RXZ2x1WEhmT0ZBcDZCTHNvalkvTTBETDN1enlOT0JXb2h4cTJmZjRmMzJrQW1SdkIxYWVWdlBXeVk5aFhjcUN2VldlOThZUTU3dXNkQ0V6NVVkdHFYTWg4OXU3eU9hcENKcHNUNnhaUGlSclhGRHRJNUM3a2F3RXRRbFlPM25FYThtVWNOYlJ0TXVIeWVkNXRDS3dtdVJHTkpKeWJXNU9mSVpnMng4U1JUUHBhalp6MG10dzdJRjZsNzI2d1N4VUM4ZHpTRGh6ZHROTEV1YTd3YWZiblVRc3piVFpaMWxZVW9PMGxDaTNzOG5jYVRJNG5RSG53QkVpSGVpbmxzZG9jOVlveTVZcWdyVGxQRG1ObVlibDVFako5ZXgzNVUyWlZlMUlNU053SDhhRnBpYkVmRVdRMDNTbk9aM2tMNm5NM3RMU0dEQ3RNNVk3SmhEYit3TmVFZUxlTzFuVXBta3hLV2VEZ0ZoYTV5eEVINGV5dHYvUjcvVDU5WUlWSlEzc0pqS2Jta29lS29sVDUwS2pCUU5rNWptYmx4ZzNPRlM3d0J3MkxpMGhaK29rNTZVZEIycFh1bFBvNGlkUUtFV1FXa3k3SVE5amNYeHpjdndTQi95eU5vMDJmMldTMmpIRVlybzhyVTMzdDFJYk42OHdHanlSaDQyRzlQSjdqeE8yOGJ0YnVtdlRzUHVXNVZlZEc2d1dmTEhKU3AzdDlOaFhabnREYU1uSnl4NmNrMmR1czBjS0RPRmErT2tuc3lPbE9sWjdmMllud05xMXAxaUVaTXEzVXIzVHY1N3VRRVhqS2U4VzhBT3MvSHdVZDBNbHFEcHhjYzhqUlorbWtBMk54dWU1bVQwNmZ1aXJkYnRZRm5lUjhaNDRuN0MwbWxCVnJ2b1laNEwzeCt3cG1SZTV3UkowME1kUjJxdnBWT0hid1kzM3dqRlcxNmdrQVoyenVZZzQ2VlpyTzNUM2s5UFpRUFhOWktwcjZvUGZENDhqNFhUWlFuN1Z6Z1BIU09SK010K0hzcjlYQ0hlMmRZSTRTNjJSK2FZQnd5QmY5Ui9JVnplcFVUN0l0dmNYR2h6NElPNVB6d1J4Vm9Jc29ncUphWmdnMk1kdHJGNXdISjRVVzg0OXhwclRkWmpCYXQraVMxaEk4bTM0eURROEJ1UjkyZFdVNTRwM0dFNTBLL3JadVhLdXNZbHNLNU1XVjdHYjR5NGM3RWIwNDNMV243cnBBZVJvRGROblhERUZrNHU4eHpGVXdQSWFSb3F1ZHlDbnZMRDRueitQd1ZnSlMxdksrcjNtc1p5N2IxekFuc0NmbUN3N3ZxeHdqUGZ5bWV5L3ZqNWcxbEJ4MlI0U1lWZU5WOGxrMzgvVGhTeDVzNFlSNm52Z0xyVTN0dHFhV0x0ZG8wUUFma3NrNEFkb1JlL2V5U1Z2UktFVlp0dFpoZFp3TnVXYkVkL1hzZFhBQXZSZ0M4QU5NSG0vbGZUcVh4SCtxOERzenY4SFZVbHNINHNRajMwQUFBQUFTVVZPUks1Q1lJST0iPjwvaW1hZ2U+CiAgICAgICAgICAgICAgICA8L2c+CiAgICAgICAgICAgIDwvZz4KICAgICAgICA8L2c+CiAgICA8L2c+Cjwvc3ZnPg=="
        />
        <span
          style="font-family: Verdana, Geneva, sans-serif; font-weight: 600; color: white;"
          >CONNECT</span
        >
      </button>
    </form>
  </body>
</html>
`