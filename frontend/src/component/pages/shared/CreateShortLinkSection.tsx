import React, { Component, ReactElement } from 'react';

import './CreateShortLinkSection.scss';
import { TextField } from '../../form/TextField';
import { Button } from '../../ui/Button';
import { ShortLinkUsage } from './ShortLinkUsage';
import { Section } from '../../ui/Section';
import { ShortLink } from '../../../entity/ShortLink';
import { UIFactory } from '../../UIFactory';
import { validateLongLinkFormat } from '../../../validators/LongLink.validator';
import { validateCustomAliasFormat } from '../../../validators/CustomAlias.validator';
import { raiseCreateShortLinkError } from '../../../state/actions';
import { IAppState } from '../../../state/reducers';
import { Store } from 'redux';
import { GraphQLService } from '../../../service/GraphQL.service';
import { ShortLinkService } from '../../../service/ShortLink.service';
import { QrCodeService } from '../../../service/QrCode.service';

interface IProps {
  store: Store<IAppState>;
  uiFactory: UIFactory;
  graphQLService: GraphQLService;
  shortLinkService: ShortLinkService;
  qrCodeService: QrCodeService;
  onShortLinkCreated?: (shortLink: string) => void;
  onAuthenticationFailed?: () => void;
}

interface IState {
  longLink: string;
  username: string;
  room: string;
  alias: string;
  inputError?: string;
  isShortLinkPublic?: boolean;
  shouldShowUsage: boolean;
  createdShortLink: string;
  createdLongLink: string;
  qrCodeURL: string;
  description: string | ReactElement;
  valid: null | boolean;
  club: string;
  link: string;
  status: string;
}

export class CreateShortLinkSection extends Component<IProps, IState> {
  private shortLinkTextField = React.createRef<TextField>();

  constructor(props: IProps) {
    super(props);
    this.state = {
      inputError: '',
      alias: '',
      longLink: '',
      username: '',
      room: '',
      shouldShowUsage: false,
      createdShortLink: '',
      createdLongLink: '',
      qrCodeURL: '',
      description: (
        <>
          Enter the super-secret code<span aria-hidden> 🤓</span>
        </>
      ),
      valid: false,
      club: 'green',
      link: '',
      status: ''
    };
  }

  render(): ReactElement {
    return (
      <Section title={''}>
        <div className={'control create-short-link'}>
          <h1 aria-label="clublink">
            <span className={this.state.club}>club</span>
            <span className={this.state.link}>l</span>
            <span className={this.state.club}>.</span>
            <span className={this.state.link}>ink </span>
            <span className={'sand'}>/</span>
          </h1>
          <div className={'code-field-wrapper'}>
            <TextField
              aria="enter the short code of the link you want to visit"
              describedBy="code-description"
              id="code-input"
              className={`code${
                this.state.status === 'error' ? ' code-error' : ''
              }`}
              ref={this.shortLinkTextField}
              text={this.state.alias}
              placeHolder={'enter code'}
              onBlur={this.handleCustomAliasTextFieldBlur}
              onChange={this.handleAliasChange}
              onFocus={this.handleFocus}
            />
            {this.state.status === 'error' && (
              <span className="emoji" aria-hidden>
                😩
              </span>
            )}
            {this.state.status === 'success' && (
              <span className="emoji" aria-hidden>
                😍
              </span>
            )}
            {/* {this.state.alias && (
            <span
              role="button"
              className={'rocket-button'}
              onClick={this.handleReserveShortLinkClick}
            >
              🚀
            </span>
          )} */}
          </div>
        </div>
        <div className={'input-description'} id="code-description">
          {this.state.description}
        </div>
        <div className={'emoji-text-field-wrapper'}>
          <TextField
            className={`${this.state.inputError ? 'error' : ''}`}
            describedBy="longLink-error"
            text={this.state.longLink}
            placeHolder={
              'Now enter your fricking ridiculously long shitty link here...'
            }
            onBlur={this.handleLongLinkTextFieldBlur}
            onChange={this.handleLongLinkChange}
          />
          {this.state.inputError === undefined ? (
            <span className="emoji input-emoji hide-small" aria-hidden>
              😱
            </span>
          ) : (
            <span className="emoji input-emoji hide-small" aria-hidden>
              💩
            </span>
          )}
        </div>
        {this.state.inputError && (
          <div id="longLink-error" className={'input-error'}>
            {this.state.inputError}
          </div>
        )}
        <div className={'text-field-wrapper'}>
          <TextField
            aria="clubhouse handle"
            className="username"
            text={this.state.username}
            placeHolder={'@username'}
            onChange={this.handleUsernameChange}
          />
          <TextField
            className="room"
            text={this.state.room}
            placeHolder={'your room title'}
            onChange={this.handleRoomChange}
          />
        </div>
        <div className="create-short-link-btn">
          <a className="back" href="/" aria-label="back to home">
            back
          </a>
          <Button
            className={'publish'}
            onClick={this.handleCreateShortLinkClick}
            disabled={
              !!this.state.inputError ||
              !this.state.alias ||
              !this.state.longLink ||
              !this.state.username ||
              !this.state.room
            }
          >
            publish
          </Button>
        </div>
        {/* {this.props.uiFactory.createPreferenceTogglesSubSection({
          uiFactory: this.props.uiFactory,
          isShortLinkPublic: this.state.isShortLinkPublic,
          onPublicToggleClick: this.handlePublicToggleClick
        })} */}
        {this.state.shouldShowUsage && (
          <div className={'short-link-usage-wrapper'}>
            <ShortLinkUsage
              shortLink={this.state.createdShortLink}
              longLink={this.state.createdLongLink}
              qrCodeUrl={this.state.qrCodeURL}
            />
          </div>
        )}
      </Section>
    );
  }

  autoFillInLongLink(longLink: string) {
    if (!longLink) {
      return;
    }

    this.setState({
      longLink: longLink
    });

    const inputError = validateLongLinkFormat(longLink);
    if (inputError != null) {
      this.setState({
        inputError: inputError
      });
      return;
    }

    this.focusShortLinkTextField();
  }

  handleFocus = () => {
    if (this.state.status === '') {
      this.setState({
        description: (
          <>
            This link only lasts 24 hours<span aria-hidden> 💣</span>
          </>
        )
      });
    }
  };

  handleLongLinkTextFieldBlur = () => {
    const { longLink } = this.state;
    const err = validateLongLinkFormat(longLink);
    this.setState({
      inputError: err || undefined
    });
  };

  handleLongLinkChange = (newLongLink: string) => {
    this.setState({
      longLink: newLongLink
    });
  };

  handleRoomChange = (newRoom: string) => {
    this.setState({
      room: newRoom
    });
  };

  handleUsernameChange = (newUsername: string) => {
    this.setState({
      username: newUsername
    });
  };

  handleAliasChange = async (newAlias: string) => {
    if (newAlias === '') {
      this.setState({
        description: (
          <>
            This link only lasts 24hrs<span aria-hidden> ✌️</span>
          </>
        ),
        alias: newAlias,
        club: 'green',
        link: '',
        status: ''
      });
    } else {
      this.setState({
        alias: newAlias,
        club: 'sand',
        link: '',
        status: ''
      });
      // await this.isAliasAvailable(newAlias).then(response => {});
    }
  };

  handleCustomAliasTextFieldBlur = () => {
    const { alias } = this.state;
    const err = validateCustomAliasFormat(alias);
    if (this.state.status === '') {
      this.setState({
        description: (
          <>
            Enter the super-secret code<span aria-hidden> 🤓</span>
          </>
        )
      });
    }
  };

  // handleReserveShortLinkClick = () => {
  //   const { alias } = this.state;
  //   const shortLink: ShortLink = {
  //     longLink: '#',
  //     alias: alias || ''
  //   };
  //   this.props.shortLinkService
  //     .createShortLink(shortLink, this.state.isShortLinkPublic)
  //     .then(async (createdShortLink: ShortLink) => {
  //       const shortLink = this.props.shortLinkService.aliasToFrontendLink(
  //         createdShortLink.alias!
  //       );

  //       const qrCodeURL = await this.props.qrCodeService.newQrCode(shortLink);

  //       this.setState({
  //         createdShortLink: shortLink,
  //         qrCodeURL: qrCodeURL,
  //         shouldShowUsage: true
  //       });

  //       if (this.props.onShortLinkCreated) {
  //         this.props.onShortLinkCreated(shortLink);
  //       }
  //     })
  //     .catch(({ authenticationErr, createShortLinkErr }) => {
  //       if (authenticationErr) {
  //         if (this.props.onAuthenticationFailed) {
  //           this.props.onAuthenticationFailed();
  //         }
  //         return;
  //       }
  //       this.props.store.dispatch(
  //         raiseCreateShortLinkError(createShortLinkErr)
  //       );
  //     });
  // };

  handleCreateShortLinkClick = async () => {
    const { alias, longLink, username, room } = this.state;
    const shortLink: ShortLink = {
      longLink: longLink,
      alias: alias || '',
      username: username,
      room: room
    };
    let err = validateCustomAliasFormat(alias);
    if (err !== null) {
      let codeInput = document.getElementById('code-input');
      if (codeInput) {
        codeInput.focus();
      }
      this.setState({
        description: <>{err}</>,
        link: 'error',
        status: 'error'
      });
      return;
    }
    await this.isAliasAvailable(alias).then(response => {
      if (response) {
        this.props.shortLinkService
          .createShortLink(shortLink, this.state.isShortLinkPublic)
          .then(async (createdShortLink: ShortLink) => {
            const shortLink = this.props.shortLinkService.aliasToFrontendLink(
              createdShortLink.alias!
            );

            console.log(createdShortLink);

            const qrCodeURL = await this.props.qrCodeService.newQrCode(
              longLink
            );

            this.setState({
              createdShortLink: shortLink,
              createdLongLink: longLink,
              qrCodeURL: qrCodeURL
              // shouldShowUsage: true
            });

            if (this.props.onShortLinkCreated) {
              this.props.onShortLinkCreated(shortLink);
            }

            window.location.assign(`/a/published/?link=${createdShortLink.id}`);
          })
          .catch(({ authenticationErr, createShortLinkErr }) => {
            console.log(authenticationErr);
            console.log(createShortLinkErr);
            if (authenticationErr) {
              if (this.props.onAuthenticationFailed) {
                this.props.onAuthenticationFailed();
              }
              return;
            }
            this.props.store.dispatch(
              raiseCreateShortLinkError(createShortLinkErr)
            );
          });
      } else {
        let codeInput = document.getElementById('code-input');
        if (codeInput) {
          codeInput.focus();
        }
        this.setState({
          description: (
            <>
              Oops! <span className="error">{alias}</span> is invalid
            </>
          ),
          link: 'error',
          status: 'error'
        });
      }
    });
  };

  isAliasAvailable = async (alias: string): Promise<boolean> => {
    return await this.props.graphQLService
      .query('http://localhost:8080/graphql', {
        query: `query {
          authQuery {
            activeShortLink(alias: "${alias}") {
              alias
              longLink
              expireAt
            }
          }
        }`,
        variables: {}
      })
      .then(results => {
        this.setState({
          description: (
            <>
              Oops! <span className="error">{alias}</span> is unavailable
            </>
          ),
          link: 'error',
          status: 'error'
        });
        return false;
      })
      .catch(error => {
        // this.setState({
        //   description: (
        //     <>
        //       Oops! <span className="error">{alias}</span> is invalid
        //     </>
        //   ),
        //   link: 'error',
        //   status: 'error'
        // });
        return true;
      });
  };

  handlePublicToggleClick = (enabled: boolean) => {
    this.setState({
      isShortLinkPublic: enabled
    });
  };

  focusShortLinkTextField = () => {
    if (!this.shortLinkTextField.current) {
      return;
    }
    this.shortLinkTextField.current.focus();
  };
}
